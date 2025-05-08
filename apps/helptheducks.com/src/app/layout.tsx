import type { Metadata, Viewport } from "next";
import "./globals.css";
import { GoogleTagManager } from "@next/third-parties/google";

const description =
  "Multiple investigations have exposed Reichardt Duck Farm for rampant disease and criminal animal cruelty.";
// Note: title should be descriptive enough for sharing on Facebook.
const title = "Help the Ducks at Reichardt";
export const metadata: Metadata = {
  title,
  description,
  icons: {
    icon: "/favicon.png",
  },
  openGraph: {
    type: "website",
    images: "https://helptheducks.com/og-image.jpeg",
    url: "https://helptheducks.com",
    description,
    siteName: title,
  },
};

export const viewport: Viewport = {
  width: "device-width",
  initialScale: 1.0,
};

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
