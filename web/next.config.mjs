/** @type {import('next').NextConfig} */
const nextConfig = {
  assetPrefix: "/docs-static",
  async rewrites() {
    return {
      beforeFiles: [
        {
          source: "/docs-static/_next/:path+",
          destination: "/_next/:path+",
        },
        {
          source: "/install",
          destination:
            "https://raw.githubusercontent.com/TreasureUzoma/readmit/main/scripts/install",
        },
        {
          source: "/install.ps1",
          destination:
            "https://raw.githubusercontent.com/TreasureUzoma/readmit/main/scripts/install.ps1",
        },
      ],
    }
  },
}

export default nextConfig
