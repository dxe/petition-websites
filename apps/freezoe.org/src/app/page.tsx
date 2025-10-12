import { ScrollButton } from "@dxe/petitions-components/scroll-button";
import { Section } from "@dxe/petitions-components/section";
import Image from "next/image";

import hero from "./img/zoe-hero-2.jpg";
import criminalNeglectOfChicken from "./img/criminal-neglect-of-chicken.jpg";
import chickenZoonoticPathogens from "./img/chicken-zoonotic-pathogens.jpg";
import chickenStuckOnBackInFactoryFarm from "./img/chicken-stuck-on-back-in-factory-farm.jpg";
import falseMarketing1 from "./img/false-marketing-1.jpg";
import falseMarketing2 from "./img/false-marketing-2.jpg";
import falseMarketing6 from "./img/false-marketing-6.jpg";
import sufferingChickenInFactoryFarm from "./img/suffering-chicken-in-factory-farm.jpg";
import chickenCrowdingInFactoryFarm from "./img/chicken-crowding-in-factory-farm.jpg";
import chickensBoiledAliveSlaughterhouse1 from "./img/chickens-boiled-alive-slaughterhouse-1.jpg";
import investigatoryReportFadingScreenshot from "./img/investigatory-report-fading-screenshot.jpg";
import { cn } from "@dxe/petitions-components/utils";
import { StaticImport } from "next/dist/shared/lib/get-img-props";
import { PetitionWithSuspense } from "./petition";
import { PressHits } from "./press-hits";

export default function HomePage() {
  return (
    <div className="flex flex-col gap-6 items-center">
      <Hero />
      <PetitionSection />
      <MoreBackgroundSection />
      <KeyFindingsSection />
      <FullInvestigatoryReport />
      <PressHits />
      <MarketingVsReality />
      <Footer />
    </div>
  );
}

function Hero() {
  return (
    <section
      className="md:min-h-[90vh] w-full text-white bg-cover flex flex-col"
      style={{
        backgroundImage: `url(${hero.src})`,
        backgroundPosition: "center 30%",
        backgroundSize: "cover",
      }}
    >
      <div className="bg-black/40 w-full grow flex justify-center items-center md:px-12 py-12">
        <div className="flex flex-col gap-9 max-w-(--breakpoint-xl) md:border-l-2 p-6 w-full">
          <div className="flex flex-col gap-4">
            <h1 className="uppercase text-2xl md:text-5xl leading-[1.125] tracking-wide max-w-[16rem] md:max-w-md border-b-2 md:border-0 py-4 md:py-0">
              Free Zoe Rosenberg
            </h1>
            <p className="font-medium max-w-lg">
              Zoe was politically prosecuted. Now, her life is in danger.
            </p>
          </div>
          <ScrollButton
            className="self-start"
            variant="secondary"
            size="lg"
            scrollToId="petition-section"
          >
            Add your name
          </ScrollButton>
        </div>
      </div>
    </section>
  );
}

function PetitionSection() {
  return (
    <Section
      className="gap-12 items-center bg-slate-200 xl:rounded-lg py-12 md:px-16"
      id="petition-section"
    >
      <h2 className="font-semibold text-xl uppercase self-start text-center md:text-left w-full">
        Governor Newsom: Pardon Zoe Rosenberg
      </h2>

      <PetitionWithSuspense />
    </Section>
  );
}

function MoreBackgroundSection() {
  function ExposingPerdueYouTubeVideo() {
    return (
      <iframe
        className="relative aspect-video h-min w-full lg:w-full max-w-3xl self-center rounded-2xl"
        src="https://www.youtube.com/embed/3r4xjelwY0U?si=ksKf8aU7aArUnWkr"
        title="YouTube video player"
        frameBorder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
        referrerPolicy="strict-origin-when-cross-origin"
        allowFullScreen
      ></iframe>
    );
  }

  return (
    <Section>
      <div className="flex flex-col lg:flex-row items-center justify-evenly gap-12">
        <div className="w-full lg:w-1/2 text-left space-y-8">
          <p>
            Zoe is incarcerated in Sonoma County, CA for the "crime" of rescuing
            animals from abuse at a Perdue-owned slaughterhouse.{" "}
            <b>All along, this was a political prosecution</b>; instead of
            prosecuting Perdue for documented animal cruelty, the District
            Attorney punished Zoe for exposing it.
          </p>
          <p>
            The District Attorney charged Zoe with a ramped-up felony conspiracy
            charge to increase her maximum sentence and forced her to wear a GPS
            ankle monitor for over 20 months while awaiting trial. When trial
            finally arrived, the court severely restricted Zoe's ability to
            present evidence of the animal cruelty she had documented at Perdue
            and reported to authorities before the rescue. The jury only heard a
            small sliver of the story and Zoe was ultimately convicted and taken
            immediately into custody.
          </p>
          <p>
            <b>Her incarceration isn't just unjust; it is life-threatening.</b>{" "}
            Zoe has a serious chronic illness that requires constant medical
            attention.{" "}
            <b>
              Being deprived of her freedom and medical care while in jail could
              kill her.
            </b>
          </p>
          <p>
            We demand her immediate release.{" "}
            <b>
              Compassion should not be criminalized, and it certainly shouldn't
              be a death sentence.
            </b>
          </p>
        </div>
        <div className="w-full lg:w-1/2 flex items-center justify-center">
          <ExposingPerdueYouTubeVideo />
        </div>
      </div>
    </Section>
  );
}

function KeyFindingsSection() {
  return (
    <Section>
      <h2 className="border-b border-slate-300 pb-2 uppercase text-xl tracking-wide text-slate-800">
        Years of evidence show a consistent pattern of extreme animal cruelty.
      </h2>
      <div className="flex flex-col lg:flex-row items-center lg:items-start justify-evenly gap-12 text-center">
        <div className="flex flex-col gap-6 max-w-sm w-full">
          <Image
            src={criminalNeglectOfChicken}
            className="rounded-3xl size-96 mb-5 object-cover object-top"
            alt="Neglect of chickens"
          />
          <div className="flex flex-col gap-4">
            <div className="text-4xl uppercase font-semibold">
              Criminal Neglect
            </div>
            <p>
              Sonoma County Animal Services examined birds from a Petaluma
              Poultry factory farm in Petaluma in 2018 and identified wing and
              joint injuries, open sores, and necrotic wounds so deep that
              muscle and bone were exposed. They{" "}
              <a
                href="https://drive.google.com/file/d/1W1-z_MoRWzwtH293-hpJcfdsmHFx5riP/view?usp=sharing"
                className="underline"
              >
                referred
              </a>{" "}
              the owner of the facility as a suspect for animal cruelty to the
              Sheriff&apos;s Office, but charges were never filed.
            </p>
          </div>
        </div>
        <div className="flex flex-col gap-6 max-w-sm w-full">
          <Image
            src={chickenZoonoticPathogens}
            className="rounded-3xl size-96 mb-5 object-cover object-[63%_50%]"
            alt="Zoonotic Pathogens"
          />
          <div className="flex flex-col gap-4">
            <div className="text-4xl uppercase font-semibold">
              Zoonotic Pathogens
            </div>
            <p>
              Lab testing has identified numerous infectious pathogens and
              diseases present at Petaluma Poultry factory farms in Sonoma
              County, including Infectious Bursal Disease, Infectious Bronchitis
              Virus, and a highly antibiotic-resistant Enterococcus bacteria.
            </p>
          </div>
        </div>
        <div className="flex flex-col gap-6 max-w-sm w-full">
          <Image
            src={chickensBoiledAliveSlaughterhouse1}
            className="rounded-3xl size-96 mb-5 object-cover"
            alt="Chickens boiled alive at slaughterhouse"
          />
          <div className="flex flex-col gap-4">
            <div className="text-4xl uppercase font-semibold">
              Animals Boiled Alive
            </div>
            <p>
              Chickens are stunned and slaughtered in the dark at high speeds,
              so many of them miss the stunning stage. This means they can enter
              scalding hot water while still conscious. A whistleblower
              documented bright red chicken parts that had been discarded inside
              the Petaluma Poultry slaughterhouse, an indication that the birds
              were scalded alive and subsequently condemned.
            </p>
          </div>
        </div>
      </div>
    </Section>
  );
}

function FullInvestigatoryReport() {
  // Copied from righttorescue.com -> pages/petalumapoultry.tsx
  return (
    <a
      href="https://www.righttorescue.com/cases/petalumapoultry/Petaluma Poultry Investigatory Report 2023.pdf"
      target="_blank"
    >
      <div
        style={{
          position: "relative",
          maxWidth: "1440px",
          marginLeft: "auto",
          marginRight: "auto",
        }}
      >
        <Image
          src={investigatoryReportFadingScreenshot}
          alt="Investigatory report screenshot"
          style={{
            width: "100%",
            aspectRatio: "1920/860",
          }}
        />
        <span
          style={{
            position: "absolute",
            bottom: "10%",
            left: "50%",
            transform: "translateX(-50%)",
            fontSize: "clamp(18px, 3vw, 48px)",
            // Prevent unnecessary wrapping. I'm not sure why it's wrapping
            // without this on some smaller screen sizes, especially when
            // padding is applied which would've been useful for a
            // semi opaque background for improved visibility.
            width: "90%",
            textShadow: "#FC0 1px 0 10px",
            boxSizing: "border-box",
            color: "black",
            textAlign: "center",
          }}
        >
          Read the full investigatory report
        </span>
      </div>
    </a>
  );
}

function MarketingVsReality() {
  function Row(props: { className?: string; children: React.ReactNode }) {
    return (
      <div
        className={cn(
          "flex flex-row w-full items-start justify-evenly",
          props.className,
        )}
      >
        {props.children}
      </div>
    );
  }

  function Header(props: { className?: string; children?: React.ReactNode }) {
    return (
      <p
        className={cn(
          "text-xl lg:text-3xl self-start p-6 w-full text-center",
          props.className,
        )}
      >
        {props.children}
      </p>
    );
  }

  function CaptionedImage(props: {
    src: string | StaticImport;
    alt: string;
    className?: string;
    children?: React.ReactNode;
  }) {
    return (
      <div className="flex flex-col w-full gap-5 items-center">
        <Image
          src={props.src}
          className={cn(
            "object-cover w-full aspect-square max-w-xl lg:rounded-lg",
            props.className,
          )}
          alt={props.alt}
        />
        <div className="pl-3 pr-3 text-center">{props.children}</div>
      </div>
    );
  }

  const petalumaPoultryVideo2018 =
    "https://www.dropbox.com/scl/fi/cjfycmxdl2exa6r0opc6z/B-ROLL-CONDENSED-FOOTAGE.mp4?rlkey=vki4hocvxo0depz7se7e0lb5c&e=1&dl=0";

  return (
    <Section className="p-0">
      <h2 className="border-b border-slate-300 pb-2 pl-4 uppercase text-xl tracking-wide text-slate-800">
        Marketing vs. The reality
      </h2>
      <div
        className="flex flex-col items-stretch justify-evenly gap-10 pb-10"
        style={{
          background:
            "linear-gradient(to right, #bcceeb 0%, #bcceeb calc(50% - 10px), #d4c1be calc(50% + 10px), #d4c1be 100%)",
        }}
      >
        <Row className="sticky top-0 backdrop-brightness-75">
          <Header className="bg-[#bcceeb]/75">Marketing</Header>
          <Header className="bg-[#d4c1be]/75">Reality</Header>
        </Row>
        <Row>
          <CaptionedImage
            src={falseMarketing1}
            alt="Petaluma Poultry false marketing"
          >
            Petaluma Poultry&apos;s ROSIE brand logo
          </CaptionedImage>
          <CaptionedImage
            src={sufferingChickenInFactoryFarm}
            alt="Petaluma Poultry chicken suffering"
          >
            A baby chicken stuck on their back at a{" "}
            <a href={petalumaPoultryVideo2018} className="underline">
              Petaluma Poultry factory farm in 2018
            </a>
          </CaptionedImage>
        </Row>
        <Row>
          <CaptionedImage
            src={falseMarketing6}
            className="object-bottom"
            alt="Petaluma Poultry false marketing"
          >
            <a
              href="https://www.instagram.com/p/Cp-RKiptIju/"
              className="underline"
            >
              A Petaluma Poultry ad
            </a>{" "}
            posted on their Instagram in 2023
          </CaptionedImage>
          <CaptionedImage
            src={chickenCrowdingInFactoryFarm}
            alt="Petaluma Poultry crowding"
          >
            Crowded birds inside a{" "}
            <a href={petalumaPoultryVideo2018} className="underline">
              Petaluma Poultry factory farm in 2018
            </a>
          </CaptionedImage>
        </Row>
        <Row>
          <CaptionedImage
            src={falseMarketing2}
            className="object-bottom"
            alt="Petaluma Poultry false marketing"
          >
            <a
              href="https://www.instagram.com/p/DBl5W1cKxXl/"
              className="underline"
            >
              A Petaluma Poultry ad
            </a>{" "}
            posted on their Instagram in 2024
          </CaptionedImage>
          <CaptionedImage
            src={chickenStuckOnBackInFactoryFarm}
            alt="Petaluma Poultry chicken stuck on back"
          >
            A sick chicken stuck on her back at a{" "}
            <a
              href="https://media.dxe.io/wl/?id=Vrgs7DjwG4Zl9GCnzurcX5Jxkt5uj85I&fmode=open"
              className="underline"
            >
              Petaluma Poultry factory farm in 2023
            </a>
          </CaptionedImage>
        </Row>
      </div>
    </Section>
  );
}

function Footer() {
  return (
    <Section className="text-center pt-4 pb-12 text-sm md:border-t border-slate-300 text-slate-600">
      &copy; {new Date().getFullYear()} Help The Chickens
    </Section>
  );
}
