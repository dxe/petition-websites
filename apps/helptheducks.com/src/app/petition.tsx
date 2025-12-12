"use client";

import { useSearchParams } from "next/navigation";
import { EmailPetition } from "@dxe/email-petition/email-petition";
import { DEFAULT_MESSAGE } from "@/data/petition-message";
import { Suspense } from "react";

export function PetitionWithSuspense() {
  return (
    <Suspense>
      {/* Suspense is Required for useSearchParams.*/}
      <Petition></Petition>
    </Suspense>
  );
}

export function Petition() {
  const searchParams = useSearchParams();

  const thermometerStartDate = process.env.NEXT_PUBLIC_THERMOMETER_START_DATE;
  const thermometerGoalEnv = process.env.NEXT_PUBLIC_THERMOMETER_GOAL;
  const thermometerOffsetEnv = process.env.NEXT_PUBLIC_THERMOMETER_OFFSET;

  const thermometerProps = thermometerStartDate
    ? {
        thermometerStartDate,
        thermometerGoal: thermometerGoalEnv
          ? parseInt(thermometerGoalEnv, 10)
          : undefined,
        thermometerOffset: thermometerOffsetEnv
          ? parseInt(thermometerOffsetEnv, 10)
          : undefined,
      }
    : {};

  return (
    <EmailPetition
      petitionId={process.env.NEXT_PUBLIC_PETITION_ID!}
      campaignName={process.env.NEXT_PUBLIC_CAMPAIGN_NAME!}
      defaultMessage={DEFAULT_MESSAGE}
      onSubmit={onSubmit}
      debug={searchParams.get("debug") === "true"}
      test={searchParams.get("test") === "true"}
      {...thermometerProps}
    />
  );
}

function onSubmit() {
  window.dataLayer?.push({
    event: "form_submitted",
  });
}
