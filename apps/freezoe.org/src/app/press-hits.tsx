"use client";

import { Section } from "@dxe/petitions-components/section";
import { useEffect, useState } from "react";

export function PressHits() {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  return (
    <Section>
      <h2 className="border-b border-slate-300 pb-2 uppercase text-xl tracking-wide text-slate-800">
        Perdue&apos;s Petaluma Poultry in the News
      </h2>

      <div className="flex flex-col lg:flex-row items-center lg:items-start justify-evenly gap-6">
        <div className="flex flex-col gap-6 max-w-96 w-full">
          {/* The Press Democrat: Perdue&apos;s Petaluma poultry plant struggles to eliminate bacteria that can make people sick */}
          <div className="iframely-embed">
            <div
              className="iframely-responsive"
              style={{ paddingBottom: "59.375%", paddingTop: "120px" }}
            >
              <a
                href="https://www.pressdemocrat.com/article/news/perdues-petaluma-poultry-plant-struggles-to-limit-pathogens/"
                data-iframely-url="//iframely.net/iNsiACG"
              ></a>
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-6 max-w-96 w-full">
          {/* The Intercept: Dangerous Pathogens and Cruelty Law Violations at Perdue Subsidiary, Animal Rights Report Alleges */}
          <div className="iframely-embed">
            <div
              className="iframely-responsive"
              style={{ paddingBottom: "59.375%", paddingTop: "120px" }}
            >
              <a
                href="https://production.public.theintercept.com/2023/06/13/perdue-chicken-slaughterhouse-animal-cruelty-dxe/"
                data-iframely-url="//iframely.net/bA6f35V"
              ></a>
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-6 max-w-96 w-full">
          {/* The San Francisco Examiner: Chef Tyler Florence distances from poultry farm and animal-rights heat */}
          <div className="iframely-embed">
            <div
              className="iframely-responsive"
              style={{ paddingBottom: "59.375%", paddingTop: "120px" }}
            >
              <a
                href="https://www.sfexaminer.com/news/business/chef-tyler-florence-distances-from-poultry-farm-and-animal-rights-heat/article_64da5ab4-601a-11ef-a358-9774850aac1b.html"
                data-iframely-url="//iframely.net/qMrpLTW"
              ></a>
            </div>
          </div>
        </div>
      </div>
      {
        // Wait for hydration before loading of the Iframely embed script
        // to avoid hydration errors.
        isClient && <script async src="//iframely.net/embed.js"></script>
      }
    </Section>
  );
}
