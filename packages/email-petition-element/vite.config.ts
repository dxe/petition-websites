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
