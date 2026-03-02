"use client"

import { ReactNode } from "react"
import dynamic from "next/dynamic"

const Analytics = dynamic(() => import("supametrics"), {
  ssr: false,
})

export default function Template({ children }: { children: ReactNode }) {
  return (
    <div>
      {children}
      <Analytics
        url="https://supametrics-go-server.onrender.com"
        client="supm_1b07e65cb713f48fffccc70fbcbc8013"
      />
    </div>
  )
}
