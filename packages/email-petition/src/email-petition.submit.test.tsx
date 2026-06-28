import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { http, HttpResponse } from "msw";
import { EmailPetition, MailerPayload } from "./email-petition";
import { server } from "./test/server";
import { TEST_CAPTCHA_TOKEN } from "./test/setup";

// These URLs are derived from the same env vars the component reads (set in
// ./test/setup), so handlers and component stay in sync automatically.
const PETITION_URL = `${process.env.NEXT_PUBLIC_PETITIONS_API_ROOT}/sign`;
const MAILER_URL = `${process.env.NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT}/message/create`;

/**
 * Installs handlers for both endpoints that capture the outgoing request and
 * return success. Returns getters for the parsed bodies so a test can assert on
 * exactly what the component sent.
 */
function captureBothEndpoints() {
  const captured: {
    petition: Record<string, string> | null;
    petitionContentType: string | null;
    mailer: MailerPayload | null;
    mailerContentType: string | null;
  } = {
    petition: null,
    petitionContentType: null,
    mailer: null,
    mailerContentType: null,
  };

  server.use(
    http.post(PETITION_URL, async ({ request }) => {
      // Petition API is sent as application/x-www-form-urlencoded.
      captured.petitionContentType = request.headers.get("content-type");
      const params = new URLSearchParams(await request.text());
      captured.petition = Object.fromEntries(params.entries());
      return new HttpResponse(null, { status: 200 });
    }),
    http.post(MAILER_URL, async ({ request }) => {
      // Mailer API is sent as JSON.
      captured.mailerContentType = request.headers.get("content-type");
      captured.mailer = (await request.json()) as MailerPayload;
      return HttpResponse.json({ ok: true });
    }),
  );

  return captured;
}

/** Types the fields common to both location modes. */
async function fillContactFields(user: ReturnType<typeof userEvent.setup>) {
  await user.type(screen.getByLabelText(/full name/i), "Jane Doe");
  await user.type(screen.getByLabelText(/email/i), "jane@example.com");
  await user.type(screen.getByLabelText(/phone/i), "5551234567");
}

describe("EmailPetition submission — sfOnly mode", () => {
  it("sends the entered details (and fixed SF location) to both endpoints", async () => {
    const user = userEvent.setup();
    const captured = captureBothEndpoints();

    render(
      <EmailPetition
        petitionId="my-petition"
        campaignName="my-campaign"
        defaultMessage="I am [Your name] from [Your city] and I support this."
        locationInputMode="sfOnly"
        debug={false}
        test={false}
      />,
    );

    await fillContactFields(user);
    await user.click(
      screen.getByRole("checkbox", { name: /resident of san francisco/i }),
    );
    await user.click(screen.getByRole("button", { name: /submit/i }));

    await waitFor(() => expect(captured.mailer).not.toBeNull());

    expect(captured.petitionContentType).toContain(
      "application/x-www-form-urlencoded",
    );
    expect(captured.mailerContentType).toContain("application/json");

    expect(captured.petition).toEqual({
      id: "my-petition",
      name: "Jane Doe",
      email: "jane@example.com",
      phone: "5551234567",
      city: "san francisco",
      state: "ca",
      country: "us",
      fullHref: window.location.href,
    });

    expect(captured.mailer).toEqual({
      name: "Jane Doe",
      email: "jane@example.com",
      phone: "5551234567",
      outside_us: false,
      city: "san francisco",
      state: "ca",
      country: "us",
      message: "I am Jane Doe from san francisco and I support this.",
      campaign: "my-campaign",
      token: TEST_CAPTCHA_TOKEN,
    });
  });

  it("prefixes test ids/campaign when test mode is on", async () => {
    const user = userEvent.setup();
    const captured = captureBothEndpoints();

    render(
      <EmailPetition
        petitionId="my-petition"
        campaignName="my-campaign"
        defaultMessage="A sufficiently long message body."
        locationInputMode="sfOnly"
        debug={false}
        test={true}
      />,
    );

    await fillContactFields(user);
    await user.click(
      screen.getByRole("checkbox", { name: /resident of san francisco/i }),
    );
    await user.click(screen.getByRole("button", { name: /submit/i }));

    await waitFor(() => expect(captured.mailer).not.toBeNull());

    expect(captured.petition?.id).toBe("test:my-petition");
    expect(captured.mailer?.campaign).toBe("test:my-campaign");
  });
});

describe("EmailPetition submission — zipWithSonomaCountyCity mode", () => {
  it("sends zip, auto-selected Sonoma city, and US country to both endpoints", async () => {
    const user = userEvent.setup();
    const captured = captureBothEndpoints();

    render(
      <EmailPetition
        petitionId="sonoma-petition"
        campaignName="sonoma-campaign"
        defaultMessage="I am [Your name] from [Your city] and I support this."
        locationInputMode="zipWithSonomaCountyCity"
        debug={false}
        test={false}
      />,
    );

    await fillContactFields(user);
    // 95401 maps to a single city (Santa Rosa), which the component auto-selects
    // into react-hook-form state (the Radix trigger doesn't render the text while
    // closed, so we assert the selection via the captured request body instead).
    await user.type(screen.getByLabelText(/zip code/i), "95401");

    await user.click(screen.getByRole("button", { name: /submit/i }));

    await waitFor(() => expect(captured.mailer).not.toBeNull());

    expect(captured.petitionContentType).toContain(
      "application/x-www-form-urlencoded",
    );
    expect(captured.mailerContentType).toContain("application/json");

    expect(captured.petition).toEqual({
      id: "sonoma-petition",
      name: "Jane Doe",
      email: "jane@example.com",
      phone: "5551234567",
      zip: "95401",
      city: "Santa Rosa",
      country: "United States",
      fullHref: window.location.href,
    });

    expect(captured.mailer).toEqual({
      name: "Jane Doe",
      email: "jane@example.com",
      phone: "5551234567",
      outside_us: false,
      zip: "95401",
      city: "Santa Rosa",
      country: "United States",
      message: "I am Jane Doe from Santa Rosa and I support this.",
      campaign: "sonoma-campaign",
      token: TEST_CAPTCHA_TOKEN,
    });
  });

  it("omits location and sends outside_us:true when signer is outside the US", async () => {
    const user = userEvent.setup();
    const captured = captureBothEndpoints();

    render(
      <EmailPetition
        petitionId="intl-petition"
        campaignName="intl-campaign"
        defaultMessage="I, [Your name], stand against this."
        locationInputMode="zipWithSonomaCountyCity"
        debug={false}
        test={false}
      />,
    );

    await fillContactFields(user);
    await user.click(
      screen.getByRole("checkbox", { name: /outside the united states/i }),
    );
    await user.click(screen.getByRole("button", { name: /submit/i }));

    await waitFor(() => expect(captured.mailer).not.toBeNull());

    // No zip/city/country keys when outside the US.
    expect(captured.petition).toEqual({
      id: "intl-petition",
      name: "Jane Doe",
      email: "jane@example.com",
      phone: "5551234567",
      fullHref: window.location.href,
    });

    expect(captured.mailer).toEqual({
      name: "Jane Doe",
      email: "jane@example.com",
      phone: "5551234567",
      outside_us: true,
      message: "I, Jane Doe, stand against this.",
      campaign: "intl-campaign",
      token: TEST_CAPTCHA_TOKEN,
    });
  });
});

describe("EmailPetition submission — sequential-call contract", () => {
  // The component intentionally awaits the petition call before the mailer call
  // so a failed first request never triggers a duplicate email on retry.
  // The failed submit rethrows out of react-hook-form's handler as a floating
  // rejection; swallow the expected one so it doesn't fail the run.
  const swallowExpectedRejection = (reason: unknown) => {
    if (
      reason instanceof Error &&
      reason.message === "Error submitting message"
    )
      return;
    throw reason;
  };

  beforeEach(() => {
    vi.spyOn(window, "alert").mockImplementation(() => {});
    vi.spyOn(console, "error").mockImplementation(() => {});
    process.on("unhandledRejection", swallowExpectedRejection);
  });

  afterEach(() => {
    process.off("unhandledRejection", swallowExpectedRejection);
    vi.restoreAllMocks();
  });

  it("does not call the mailer when the petition request fails", async () => {
    const user = userEvent.setup();
    let mailerCalled = false;

    server.use(
      http.post(PETITION_URL, () => new HttpResponse(null, { status: 500 })),
      http.post(MAILER_URL, () => {
        mailerCalled = true;
        return HttpResponse.json({ ok: true });
      }),
    );

    render(
      <EmailPetition
        petitionId="my-petition"
        campaignName="my-campaign"
        defaultMessage="A sufficiently long message body."
        locationInputMode="sfOnly"
        debug={false}
        test={false}
      />,
    );

    await fillContactFields(user);
    await user.click(
      screen.getByRole("checkbox", { name: /resident of san francisco/i }),
    );
    await user.click(screen.getByRole("button", { name: /submit/i }));

    // The component alerts the user on the petition failure.
    await waitFor(() => expect(window.alert).toHaveBeenCalled());
    expect(mailerCalled).toBe(false);
  });
});
