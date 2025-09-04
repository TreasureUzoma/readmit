import { NextRequest, NextResponse } from "next/server"
import { createClient } from "@supabase/supabase-js"

import { supabaseAnonKey, supabaseUrl } from "@/config/supabase-config"
import { genai } from "@/lib/ai"

export async function POST(req: NextRequest) {
  try {
    const { fileName, mode } = await req.json()

    if (!fileName) {
      return NextResponse.json(
        { error: "fileName is required" },
        { status: 400 }
      )
    }
    if (!mode) {
      return NextResponse.json({ error: "mode is required" }, { status: 400 })
    }

    const supabase = createClient(supabaseUrl, supabaseAnonKey)

    const { data, error } = await supabase.storage
      .from("codebases")
      .createSignedUrl(fileName, 60)

    if (error) {
      console.error("Supabase createSignedUrl error:", error)
      return NextResponse.json(
        { error: "Failed to create signed URL" },
        { status: 500 }
      )
    }

    const signedUrl = data.signedUrl
    console.log("Generated signed URL:", signedUrl)

    const fileResponse = await fetch(signedUrl)
    const fileContent = await fileResponse.text()

    const result = await genai({ mode, input: fileContent })
    console.log(result)

    return NextResponse.json({
      success: true,
      fileUrl: signedUrl,
      [mode]: result,
    })
  } catch {
    console.error("API handler error:")
    return NextResponse.json(
      { error: "Internal Server Error" },
      { status: 500 }
    )
  }
}
