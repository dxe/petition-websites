import { z } from "zod";
import { SonomaCities } from "./zipcodes";

const EmptyStringToUndefined = z.literal("").transform(() => undefined);

export function makePetitionFormSchema(options?: {
  citiesByZip?: Record<string, string[]>;
}) {
  const citiesByZip = options?.citiesByZip;

  return z
    .object({
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
      message: z
        .string()
        .min(10, { message: "Message must be at least 10 characters" })
        .max(10_000, { message: "Please limit message to 10,000 characters" }),
    })
    .superRefine((data, ctx) => {
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
      const citiesForZip = data.zip ? citiesByZip?.[data.zip] : undefined;
      const isInCityScope = !!(citiesForZip && citiesForZip.length);

      // If outside of the scoped area, throw away the city.
      if (!isInCityScope) {
        data.city = undefined;
      }

      // If inside the scoped area, city is required.
      if (isInCityScope && !data.city) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "City is required",
          path: ["city"],
        });
      }
    });
}

export const PetitionFormSchema = makePetitionFormSchema({
  citiesByZip: SonomaCities,
});

export type PetitionForm = z.infer<ReturnType<typeof makePetitionFormSchema>>;
