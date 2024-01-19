import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import { TanStackRouterVite } from "@tanstack/router-vite-plugin";
import path from "path";
import { VitePluginRadar } from "vite-plugin-radar";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    TanStackRouterVite(),
    VitePluginRadar({
      /**
       * enable or disable scripts injection in development
       * default: false
       */
      enableDev: true,
      // Google Analytics (multiple tag can be set with an array)
      analytics: [
        {
          id: "G-2WJVQ0EX4G",
          disable: false,
          config: {
            cookie_domain: "auto",
            cookie_expires: 63072000,
            cookie_prefix: "none",
            cookie_update: true,
            cookie_flags: "",
            send_page_view: true,
            allow_google_signals: true,
            allow_ad_personalization_signals: true,
          },
        },
      ],
    }),
  ],
  resolve: {
    alias: {
      "~": path.resolve(__dirname, "./src"),
    },
  },
});
