import { fileURLToPath } from "node:url";
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

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
    // Styles are imported with `?inline` in src/index.tsx and rendered into the
    // shadow root, so no document-level CSS injection plugin is needed.
    plugins: [tailwindcss(), react()],
    define: {
      // React reads process.env.NODE_ENV for its dev/prod branches. Vite's
      // library mode leaves it untouched (it assumes a downstream bundler), but
      // this IIFE is loaded directly in the browser, so we must replace it here
      // or `process` is undefined at runtime.
      "process.env.NODE_ENV": JSON.stringify(
        isDev ? "development" : "production",
      ),
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
      ? {
          // Serve the dev demo page (dev-preview/index.html) at `/`.
          //
          // Because the dev root is dev-preview/, the page can't reach the
          // component source with a relative path (it'd be clamped at the root),
          // so it imports `@src/index.tsx` and we alias `@src` back to ../src.
          root: "dev-preview",
          resolve: {
            alias: {
              "@src": fileURLToPath(new URL("./src", import.meta.url)),
            },
          },
        }
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
