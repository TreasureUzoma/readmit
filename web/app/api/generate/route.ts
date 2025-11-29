/* eslint-disable @typescript-eslint/no-explicit-any */

import { NextRequest, NextResponse } from "next/server"
import { createClient } from "@supabase/supabase-js"

import { supabaseServiceRoleKey, supabaseUrl } from "@/config/supabase-config"
import { genai } from "@/lib/ai"

export async function POST(req: NextRequest) {
  try {
    const { fileName, mode } = await req.json()

    if (!fileName)
      return NextResponse.json(
        { error: "fileName is required" },
        { status: 400 }
      )
    if (!mode)
      return NextResponse.json({ error: "mode is required" }, { status: 400 })

    const supabase = createClient(supabaseUrl, supabaseServiceRoleKey)

    // Generate signed URL
    const { data, error } = await supabase.storage
      .from("codebases")
      .createSignedUrl(fileName, 60)

    if (error || !data?.signedUrl) {
      console.error("Supabase createSignedUrl error:", error)
      return NextResponse.json(
        { error: "Failed to create signed URL" },
        { status: 500 }
      )
    }

    const signedUrl = data.signedUrl
    console.log("Generated signed URL:", signedUrl)

    // Fetch file content
    const fileResponse = await fetch(signedUrl)
    const fileContent = await fileResponse.text()

    // Send to AI
    const result = await genai({ mode, input: fileContent })
    console.log(result)

    // List files in bucket
    const { data: fileList } = await supabase.storage
      .from("codebases")
      .list("", { limit: 20 })

    let deleteError: any = null

    if (fileList && fileList.length > 0) {
      const fileNames = fileList.map((file: any) => file.name)

      const { data: deleteData, error } = await supabase.storage
        .from("codebases")
        .remove(fileNames)

      deleteError = error
      console.log(
        "Detected files from Supabase storage:",
        fileNames,
        deleteData,
        deleteError
      )
      console.log("Deleted file from Supabase storage:", fileNames)
    }

    if (deleteError) console.error("Supabase file deletion error:", deleteError)

    return NextResponse.json({
      success: true,
      fileUrl: signedUrl,
      [mode]: result,
    })
  } catch (err) {
    console.error("API handler error:", err)
    return NextResponse.json(
      { error: "Internal Server Error" },
      { status: 500 }
    )
  }
}
