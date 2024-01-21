import {
  DEFAULT_MESSAGE,
  PetitionForm,
  PetitionFormSchema,
} from "../data/petition.ts";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form.tsx";
import { Input } from "./ui/input.tsx";
import { Button } from "./ui/button.tsx";
import { Textarea } from "./ui/textarea.tsx";
import { Checkbox } from "./ui/checkbox.tsx";
import { SonomaCities } from "../data/zipcodes.ts";
import { cn } from "~/utils.ts";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select.tsx";
import ky from "ky";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert.tsx";
import { LoaderIcon, MailCheckIcon } from "lucide-react";
import ReCAPTCHA from "react-google-recaptcha";

const PETITION_API_URL = "https://petitions-229503.appspot.com/api/sign";

const CAMPAIGN_MAILER_API_URL = `${import.meta.env.PROD ? "https://helptheducks.dxe.io" : "http://localhost:3333"}/message/create`;

declare global {
  interface Window {
    dataLayer?: unknown[];
  }
}

const CAPTCHA_SITE_KEY = "6LdiglcpAAAAAM9XE_TNnAiZ22NR9nSRxHMOFn8E";

export const Petition = () => {
  const form = useForm<PetitionForm>({
    resolver: zodResolver(PetitionFormSchema),
    defaultValues: {
      name: "",
      email: "",
      phone: "",
      outsideUS: false,
      zip: "",
      city: "",
      message: DEFAULT_MESSAGE,
    },
  });
  const {
    formState: { dirtyFields },
    setValue,
    getValues,
    watch,
    handleSubmit,
    control,
    resetField,
  } = form;
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSubmitted, setIsSubmitted] = useState(false);

  const onSubmit = useMemo(
    () =>
      handleSubmit(async (data) => {
        window.dataLayer?.push({
          event: "form_submitted",
        });
        setIsSubmitting(true);
        if (!recaptchaRef.current) {
          alert("Error loading captcha. Please refresh the page & try again.");
          setIsSubmitting(false);
          return;
        }
        const token = await recaptchaRef.current.executeAsync();
        if (!token) {
          alert("Captcha error. Please refresh the page & try again.");
          setIsSubmitting(false);
          return;
        }
        // We purposefully do these one at a time. If the first one fails,
        // we don't want to submit the second one. This allows the user to
        // resubmit the form without causing duplicate emails to be sent.
        const petitionResp = await ky.post(PETITION_API_URL, {
          body: new URLSearchParams({
            id: "helptheducks",
            name: data.name,
            email: data.email,
            ...(data.phone && { phone: data.phone }),
            ...(data.zip && { zip: data.zip }),
            ...(data.city && { city: data.city }),
            ...(!data.outsideUS && { country: "United States" }),
            fullHref: window.location.href,
          }),
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          throwHttpErrors: false,
        });
        if (petitionResp.status !== 200) {
          setIsSubmitting(false);
          alert("Error submitting. Please try again.");
          throw new Error("Error submitting petition");
        }
        const campaignMailerResp = await ky.post(CAMPAIGN_MAILER_API_URL, {
          json: {
            name: data.name,
            email: data.email,
            ...(data.phone && { phone: data.phone }),
            outside_us: data.outsideUS,
            ...(data.zip && { zip: data.zip }),
            ...(data.city && { city: data.city }),
            message: data.message,
            token,
          },
          headers: {
            "Content-Type": "application/json",
          },
          throwHttpErrors: false,
        });
        if (campaignMailerResp.status !== 200) {
          setIsSubmitting(false);
          alert("Error submitting. Please try again.");
          throw new Error("Error submitting message");
        }
        setIsSubmitted(true);
        setIsSubmitting(false);
      }),
    [handleSubmit],
  );

  const outsideUS = watch("outsideUS");
  const zip = watch("zip");
  const isInSonomaCounty = useMemo(() => {
    return zip && zip in SonomaCities;
  }, [zip]);
  const cities = useMemo(() => {
    if (!isInSonomaCounty) {
      return [];
    }
    return SonomaCities[zip as keyof typeof SonomaCities];
  }, [isInSonomaCounty, zip]);

  // When cities change, just select it if there's only one. Else, reset the city.
  useEffect(() => {
    if (cities.length === 1) {
      setValue("city", cities[0]);
    } else {
      setValue("city", "");
    }
  }, [cities, setValue]);

  const injectValuesIntoMessage = useCallback(
    (name: string | undefined, city: string | undefined) => {
      if (dirtyFields.message) {
        console.log(
          "Skipped updating message with name or city since it has been customized.",
        );
        return;
      }
      resetField("message", {
        defaultValue: DEFAULT_MESSAGE.replace(
          "[Your name]",
          name || "[Your name]",
        ).replace("[Your city if you live in Sonoma County]", city || ""),
      });
    },
    [dirtyFields.message, resetField],
  );

  const recaptchaRef = useRef<ReCAPTCHA>(null);

  return isSubmitted ? (
    <Alert className="self-center w-fit bg-slate-100">
      <MailCheckIcon className="h-4 w-4" />
      <AlertTitle>Thank you</AlertTitle>
      <AlertDescription>
        Your message has been submitted. Thank you for taking action!
      </AlertDescription>
    </Alert>
  ) : (
    <Form {...form}>
      <form
        onSubmit={onSubmit}
        className="w-full flex flex-col md:flex-row gap-8 justify-center"
      >
        <div className="flex flex-col gap-4 basis-1/3">
          <FormField
            control={control}
            name="name"
            disabled={isSubmitting}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Full Name</FormLabel>
                <FormControl>
                  <Input
                    placeholder="Jane Doe"
                    type="text"
                    {...field}
                    onBlur={() => {
                      field.onBlur();
                      injectValuesIntoMessage(field.value, getValues("city"));
                    }}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={control}
            name="email"
            disabled={isSubmitting}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    type="email"
                    placeholder="janedoe@gmail.com"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={control}
            name="phone"
            disabled={isSubmitting}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Phone Number</FormLabel>
                <FormControl>
                  <Input type="tel" placeholder="888-888-8888" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={control}
            name="zip"
            disabled={outsideUS || isSubmitting}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Zip Code</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="95401"
                    {...field}
                    onBlur={() => {
                      field.onBlur();
                      injectValuesIntoMessage(
                        getValues("name"),
                        getValues("city"),
                      );
                    }}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={control}
            name="city"
            disabled={outsideUS || isSubmitting}
            render={({ field }) => (
              <FormItem className={cn({ hidden: !cities.length })}>
                <FormLabel>City</FormLabel>
                <Select
                  onValueChange={(val) => {
                    field.onChange(val);
                    injectValuesIntoMessage(getValues("name"), val);
                  }}
                  defaultValue={field.value}
                  value={field.value}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="Select a city" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    {cities?.map((city) => (
                      <SelectItem value={city} key={city} onBlur={field.onBlur}>
                        {city}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={control}
            name="outsideUS"
            disabled={isSubmitting}
            render={({ field }) => (
              <FormItem
                className={cn("flex gap-2 items-center", {
                  hidden: cities.length,
                })}
              >
                <FormControl>
                  <Checkbox
                    checked={field.value}
                    onCheckedChange={field.onChange}
                    aria-label="Outside the United States?"
                  />
                </FormControl>
                <FormLabel className="!mt-0">
                  Outside the United States
                </FormLabel>
                <FormDescription></FormDescription>
              </FormItem>
            )}
          />
        </div>
        <div className="flex flex-col gap-4 basis-2/3">
          <FormField
            control={control}
            name="message"
            disabled={isSubmitting}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Message</FormLabel>
                <FormControl>
                  <Textarea {...field} rows={18} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" disabled={isSubmitting}>
            {isSubmitting && (
              <LoaderIcon className="mr-2 h-4 w-4 animate-spin" />
            )}
            Submit
          </Button>
          <ReCAPTCHA
            ref={recaptchaRef}
            sitekey={CAPTCHA_SITE_KEY}
            badge="bottomright"
            size="invisible"
          />
          <p className="text-xs text-center">
            By signing, you agree to receive email messages from Direct Action
            Everywhere. You may unsubscribe at any time.
          </p>
        </div>
      </form>
    </Form>
  );
};
