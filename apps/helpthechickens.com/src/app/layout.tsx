import type { Metadata, Viewport } from "next";
import "./globals.css";
import { GoogleTagManager } from "@next/third-parties/google";

const description =
  "Investigations since 2018 have exposed sick and injured animals languishing without care.";
// Note: title should be descriptive enough for sharing on Facebook.
const title = "Help Stop Perdue's Animal Abuse";
export const metadata: Metadata = {
  title,
  description,
  icons: {
    icon: "/favicon.png",
  },
  openGraph: {
    type: "website",
    images: "https://helpthechickens.com/og-image.jpg",
    url: "https://helpthechickens.com",
    description,
    siteName: title,
  },
};

export const viewport: Viewport = {
  width: "device-width",
  initialScale: 1.0,
};

const GTM_CONFIG = "GTM-K68F95SH";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <GoogleTagManager gtmId="GTM-K68F95SH" />
      </head>
      <body>{children}</body>
    </html>
  );
}
