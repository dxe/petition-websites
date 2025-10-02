"use client";

import { useScrollToId } from "@/hooks/useScrollToId";
import { Button } from "@/shadcn/components/ui/button";
import React from "react";

export function ScrollButton(props: {
  className?: string;
  variant?:
    | "link"
    | "default"
    | "destructive"
    | "outline"
    | "secondary"
    | "ghost";
  size?: "default" | "sm" | "lg" | "icon";
  scrollToId: string;
  children: React.ReactNode;
}) {
  const scrollToPetition = useScrollToId(props.scrollToId);

  return (
    <Button
      className={props.className}
      variant={props.variant}
      size={props.size}
      onClick={scrollToPetition}
    >
      {props.children}
    </Button>
  );
}
