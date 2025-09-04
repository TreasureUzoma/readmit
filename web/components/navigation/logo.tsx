import Image from "next/image"
import Link from "next/link"

import { Settings } from "@/lib/meta"

export function Logo() {
  return (
    <Link href="/" className="flex items-center gap-2.5">
      <Image
        src={Settings.siteicon}
        alt={`${Settings.title} main logo`}
        width={21}
        height={21}
        loading="lazy"
        decoding="async"
        className="invert"
      />
      <h1 className="text-md font-semibold">{Settings.title}</h1>
    </Link>
  )
}
