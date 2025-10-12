import type { Metadata, Viewport } from "next";
import "./globals.css";
import { GoogleAnalytics } from "@next/third-parties/google";

const description =
  "Investigations since 2018 have exposed sick and injured animals languishing without care.";
// Note: title should be descriptive enough for sharing on Facebook.
const title = "Free Zoe Rosenberg";
export const metadata: Metadata = {
  title,
  description,
  icons: {
    icon: "/favicon.png",
  },
  openGraph: {
    type: "website",
    images: "https://freezoe.org/og-image.jpg",
    url: "https://freezoe.org",
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
      <head></head>
      <body>
        {children}
        <GoogleAnalytics gaId="G-5ZVSGD02YN" />
      </body>
    </html>
  );
}
