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

  return (
    <EmailPetition
      petitionId={process.env.NEXT_PUBLIC_PETITION_ID!}
      campaignName={process.env.NEXT_PUBLIC_CAMPAIGN_NAME!}
      defaultMessage={DEFAULT_MESSAGE}
      onSubmit={onSubmit}
      debug={searchParams.get("debug") === "true"}
      test={searchParams.get("test") === "true"}
      signatureThermometer={{
        defaultGoal: Number(process.env.NEXT_PUBLIC_THERMOMETER_GOAL!),
      }}
    />
  );
}

function onSubmit() {
  window.dataLayer?.push({
    event: "form_submitted",
  });
}
