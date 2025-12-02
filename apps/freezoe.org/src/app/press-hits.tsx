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
        Zoe&apos;s Case in the News
      </h2>

      <div className="flex flex-col lg:flex-row items-center lg:items-start justify-evenly gap-6">
        <div className="flex flex-col gap-6 max-w-96 w-full">
          {/* Berkeleyside */}
          <div className="iframely-embed">
            <div
              className="iframely-responsive"
              style={{ paddingBottom: "59.375%", paddingTop: "120px" }}
            >
              <a
                href="https://www.theguardian.com/us-news/2025/oct/24/chicken-rescue-factory-farm"
                data-iframely-url="//iframely.net/9KbWKeY9?theme=dark"
              ></a>
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-6 max-w-96 w-full">
          {/* The New York Times */}
          <div className="iframely-embed">
            <div
              className="iframely-responsive"
              style={{ paddingBottom: "59.375%", paddingTop: "120px" }}
            >
              <a
                href="https://www.nytimes.com/2025/10/29/us/animal-rights-activist-convicted-sonoma-chickens.html"
                data-iframely-url="//iframely.net/lb0hCtNx?media=0&theme=dark"
              ></a>
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-6 max-w-96 w-full">
          {/* The Mercury News */}
          <div className="iframely-embed">
            <div
              className="iframely-responsive"
              style={{ paddingBottom: "59.375%", paddingTop: "120px" }}
            >
              <a
                href="https://www.kron4.com/news/bay-area/joaquin-phoenix-leaps-to-defense-of-student-convicted-in-chicken-rescue-case/"
                data-iframely-url="//iframely.net/b5sHtxhu?theme=dark "
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
