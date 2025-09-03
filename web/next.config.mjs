
/** @type {import('next').NextConfig} */
const nextConfig = {
  assetPrefix: "/docs-static",
  async rewrites() {
    return {
      beforeFiles: [
        {
          source: '/docs-static/_next/:path+',
          destination: '/_next/:path+',
        },
      ],
    }
  },
}

export default nextConfig
