import { z } from "zod";
import { SonomaCities } from "./zipcodes";

const EmptyStringToUndefined = z.literal("").transform(() => undefined);

/**
 * Controls which location fields the petition form collects.
 *
 * - `zipWithSonomaCountyCity`: zip code (US only) plus an auto-derived city
 *   selector for Sonoma County zips, and an "Outside the United States" toggle.
 * - `sfOnly`: hides zip and city entirely and instead requires the signer to
 *   confirm they are a resident of San Francisco, CA.
 */
export const LOCATION_INPUT_MODES = [
  "zipWithSonomaCountyCity",
  "sfOnly",
] as const;

export type LocationInputMode = (typeof LOCATION_INPUT_MODES)[number];

const PetitionFormObject = z.object({
  name: z
    .string()
    .min(2, { message: "Name too short" })
    .max(255, { message: "Name too long" }),
  email: z.string().email().max(255, { message: "Email too long" }),
  phone: z
    .string()
    .regex(/^[0-9+-]*$/, {
      message: "Phone number may only contain numbers, +, and -",
    })
    .min(10, { message: "Phone number too short" })
    .max(255, { message: "Phone number too long" })
    .or(EmptyStringToUndefined)
    .optional(),
  outsideUS: z.boolean(),
  zip: z
    .string()
    .regex(/^[0-9]*$/, {
      message:
        "Zip code must only contain numbers, or be empty if 'Outside the United States' is checked.",
    })
    .length(5, {
      message:
        "Zip code must be 5 digits, or empty if 'Outside the United States' is checked.",
    })
    .or(EmptyStringToUndefined)
    .optional(),
  city: z
    .string()
    .max(255, { message: "City too long" })
    .or(EmptyStringToUndefined)
    .optional(),
  // Only used in the `sfOnly` location input mode.
  sfResident: z.boolean(),
  message: z
    .string()
    .min(10, { message: "Message must be at least 10 characters" })
    .max(10_000, { message: "Please limit message to 10,000 characters" }),
});

export function makePetitionFormSchema(locationInputMode: LocationInputMode) {
  return PetitionFormObject.superRefine((data, ctx) => {
    if (locationInputMode === "sfOnly") {
      if (!data.sfResident) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must be a resident of San Francisco, CA to sign.",
          path: ["sfResident"],
        });
      }
      return;
    }

    // zipWithSonomaCountyCity mode:
    // If outside of US, throw away the zip code.
    if (data.outsideUS) {
      data.zip = undefined;
    } else {
      // If inside US, zip code is required.
      if (!data.zip) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Zip code is required if in the United States",
          path: ["zip"],
        });
      }
    }
    const isInSonoma = !!(
      data.zip &&
      data.zip in SonomaCities &&
      SonomaCities[data.zip as keyof typeof SonomaCities]
    );
    // If outside of Sonoma, throw away the city.
    if (!isInSonoma) {
      data.city = undefined;
    }
    // If in Sonoma, city is required.
    if (isInSonoma && !data.city) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "City is required",
        path: ["city"],
      });
    }
  });
}

export type PetitionForm = z.infer<typeof PetitionFormObject>;

// Location submitted for signers in the `sfOnly` location input mode.
const SAN_FRANCISCO_LOCATION = {
  city: "san francisco",
  state: "ca",
  country: "us",
} as const;

/** Form fields the location resolvers derive a signer's location from. */
type LocationInputData = Pick<PetitionForm, "zip" | "city" | "outsideUS">;

/** Location values for a signer, resolved from their input mode. */
export type ResolvedLocation = {
  zip?: string;
  city?: string;
  state?: string;
  country?: string;
  outsideUS: boolean;
};

/**
 * Source of truth for how each {@link LocationInputMode} turns signer
 * input into a location.
 */
const LOCATION_RESOLVERS: Record<
  LocationInputMode,
  (data: LocationInputData) => ResolvedLocation
> = {
  sfOnly: () => ({
    ...SAN_FRANCISCO_LOCATION,
    outsideUS: false,
  }),
  zipWithSonomaCountyCity: (data) => ({
    ...(data.zip && { zip: data.zip }),
    ...(data.city && { city: data.city }),
    ...(!data.outsideUS && { country: "United States" }),
    outsideUS: data.outsideUS,
  }),
};

/** Resolves the location values for the given input mode. */
export function resolveLocation(
  mode: LocationInputMode,
  data: LocationInputData,
): ResolvedLocation {
  return LOCATION_RESOLVERS[mode](data);
}
