import {
  DEFAULT_MESSAGE,
  PetitionForm,
  PetitionFormSchema,
} from "../data/petition.ts";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect, useMemo } from "react";
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
  } = form;

  const onSubmit = useMemo(
    () =>
      handleSubmit((data) => {
        // TODO: submit to petition service & new backend for sending email to district attorney.
        console.log(data);
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
        alert(
          "You've already customized the message, so it couldn't be automatically updated with your new name or city. Please be sure to double check your message before submitting.",
        );
        return;
      }
      setValue(
        "message",
        DEFAULT_MESSAGE.replace("[Your name]", name || "[Your name]").replace(
          "[Your city if you live in Sonoma County]",
          city || "",
        ),
      );
    },
    [dirtyFields.message, setValue],
  );

  return (
    <Form {...form}>
      <form
        onSubmit={onSubmit}
        className="w-full flex flex-col md:flex-row gap-8 justify-center"
      >
        <div className="flex flex-col gap-4 basis-1/3">
          <FormField
            control={control}
            name="name"
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
            disabled={outsideUS}
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
          <Button type="submit">Submit</Button>
          <p className="text-xs text-center">
            By signing, you agree to receive email messages from Direct Action
            Everywhere. You may unsubscribe at any time.
          </p>
        </div>
      </form>
    </Form>
  );
};
