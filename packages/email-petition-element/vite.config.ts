import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import cssInjectedByJsPlugin from "vite-plugin-css-injected-by-js";

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "");
  const isDev = mode === "development";

  const requireEnv = (name: string): string => {
    const value = env[name];
    if (!value) {
      throw new Error(`Missing required environment variable "${name}".`);
    }
    return value;
  };

  return {
    plugins: [
      tailwindcss(),
      react(),
      ...(isDev ? [] : [cssInjectedByJsPlugin()]),
    ],
    define: {
      // The email-petition package uses Next.js-style env vars at module scope.
      // These are replaced at build time via the VITE_* equivalents in .env.
      "process.env.NEXT_PUBLIC_PETITIONS_API_ROOT": JSON.stringify(
        requireEnv("VITE_PETITIONS_API_ROOT"),
      ),
      "process.env.NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT": JSON.stringify(
        requireEnv("VITE_CAMPAIGN_MAILER_API_ROOT"),
      ),
    },
    ...(isDev
      ? {}
      : {
          build: {
            lib: {
              entry: "src/index.tsx",
              name: "EmailPetition",
              formats: ["iife"],
              fileName: () => "email-petition.js",
            },
          },
        }),
  };
});
