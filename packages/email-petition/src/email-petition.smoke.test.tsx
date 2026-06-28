import "@testing-library/jest-dom/vitest";
import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { EmailPetition } from "./email-petition";

// Smoke test: confirms the harness is wired correctly and the component renders.
describe("EmailPetition harness", () => {
  it("renders the form fields", () => {
    render(
      <EmailPetition
        petitionId="my-petition"
        campaignName="my-campaign"
        defaultMessage="Please take action on this important issue."
        locationInputMode="sfOnly"
        debug={false}
        test={false}
      />,
    );

    expect(screen.getByLabelText(/full name/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /submit/i })).toBeInTheDocument();
  });
});
