import "@testing-library/jest-dom/vitest";
import { afterAll, afterEach, beforeAll, vi } from "vitest";
import { cleanup } from "@testing-library/react";
import { server } from "./server";

// ---------------------------------------------------------------------------
// Environment variables
// ---------------------------------------------------------------------------
// email-petition.tsx reads these at module-load time to build the endpoint URLs
// (PETITION_API_URL / CAMPAIGN_MAILER_API_URL). setupFiles run before the test
// module — and therefore before the component — is imported, so setting them
// here makes the URLs deterministic. Point MSW handlers at these same origins.
process.env.NEXT_PUBLIC_PETITIONS_API_ROOT = "https://petition.test/api";
process.env.NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT = "https://mailer.test/api";

// ---------------------------------------------------------------------------
// reCAPTCHA mock
// ---------------------------------------------------------------------------
// The real <ReCAPTCHA> talks to Google and gates submit via
// recaptchaRef.current.executeAsync(). We replace it with a headless stub that
// renders nothing and resolves a fixed token, so the submit flow proceeds and
// tests can assert the token is forwarded to the mailer API.
export const TEST_CAPTCHA_TOKEN = "test-captcha-token";

vi.mock("react-google-recaptcha", async () => {
  const React = await import("react");
  const ReCAPTCHA = React.forwardRef(
    (_props: Record<string, unknown>, ref: React.Ref<unknown>) => {
      React.useImperativeHandle(ref, () => ({
        executeAsync: async () => TEST_CAPTCHA_TOKEN,
        reset: () => {},
      }));
      return null;
    },
  );
  ReCAPTCHA.displayName = "ReCAPTCHA";
  return { default: ReCAPTCHA };
});

// ---------------------------------------------------------------------------
// MSW lifecycle
// ---------------------------------------------------------------------------
// `onUnhandledRequest: "error"` makes any un-stubbed network call fail loudly,
// so a test can't silently pass while the component hits a real/unexpected URL.
beforeAll(() => server.listen({ onUnhandledRequest: "error" }));
afterEach(() => {
  cleanup();
  server.resetHandlers();
});
afterAll(() => server.close());
