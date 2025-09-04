import { NextRequest, NextResponse } from "next/server"
import { createClient } from "@supabase/supabase-js"

import { supabaseAnonKey, supabaseUrl } from "@/config/supabase-config"

const supabase = createClient(supabaseUrl, supabaseAnonKey)

export async function POST(req: NextRequest) {
  try {
    console.log("[INFO] POST /api/signed-url called")

    let body: any
    try {
      body = await req.json()
      console.log("[INFO] Request body:", body)
    } catch (jsonErr) {
      console.error("[ERROR] Failed to parse JSON body:", jsonErr)
      return NextResponse.json({ error: "Invalid JSON body" }, { status: 400 })
    }

    const { path } = body
    if (!path || typeof path !== "string") {
      console.error("[ERROR] Missing 'path' in request body")
      return NextResponse.json({ error: "Filepath required" }, { status: 400 })
    }

    console.log("[INFO] Generating signed URL for file:", path)

    // 30 min expiration
    const { data, error } = await supabase.storage
      .from("codebases")
      .createSignedUploadUrl(path)

    if (error) {
      console.error("[Supabase Error]", error)
      return NextResponse.json({ error: error.message }, { status: 500 })
    }

    console.log("[INFO] Signed URL generated:", data.signedUrl)

    return NextResponse.json({
      uploadUrl: data.signedUrl,
      path: data.path,
      token: data.token,
    })
  } catch {
    console.error("[Server Error]")
    return NextResponse.json(
      { error: "Internal Server Error" },
      { status: 500 }
    )
  }
}
