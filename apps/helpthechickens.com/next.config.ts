import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'export',
  images: {
    // Required for `output: 'export'`:
    // https://nextjs.org/docs/messages/export-image-api
    unoptimized: true
  },
};

export default nextConfig;
