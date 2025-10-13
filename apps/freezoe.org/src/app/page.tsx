import { ScrollButton } from "@dxe/petitions-components/scroll-button";
import { Section } from "@dxe/petitions-components/section";
import Image from "next/image";

import hero from "./img/zoe-hero-2.jpg";
import zoeWithAnimal from "./img/zoe-with-animal.webp";
import zoeHealthRisk from "./img/zoe-health-risk.jpg";
import zoeRescuedChicken from "./img/zoe-rescued-chickens.webp";
import { PetitionWithSuspense } from "./petition";
import { PressHits } from "./press-hits";

export default function HomePage() {
  return (
    <div className="flex flex-col gap-6 items-center">
      <Hero />
      <PetitionSection />
      <MoreBackgroundSection />
      <WhoIsZoeSection />
      <ZoesHealthSection />
      <WhyIsZoeInJailSection />
      <TakeActionSection />
      <PressHits />
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
        backgroundPosition: "center 35%",
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
  // function ExposingPerdueYouTubeVideo() {
  //   return (
  //     <iframe
  //       className="relative aspect-video h-min w-full lg:w-full max-w-3xl self-center rounded-2xl"
  //       src="https://www.youtube.com/embed/3r4xjelwY0U?si=ksKf8aU7aArUnWkr"
  //       title="YouTube video player"
  //       frameBorder="0"
  //       allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
  //       referrerPolicy="strict-origin-when-cross-origin"
  //       allowFullScreen
  //     ></iframe>
  //   );
  // }

  return (
    <Section>
      <div className="flex flex-col lg:flex-row items-center justify-evenly gap-12">
        <div className="w-full lg:w-1/2 text-left space-y-8 order-2 lg:order-1">
          <p>
            Zoe is incarcerated in Sonoma County, CA for the &quot;crime&quot;
            of rescuing animals from abuse at a Perdue-owned slaughterhouse.{" "}
            <b>All along, this was a political prosecution</b>; instead of
            prosecuting Perdue for documented animal cruelty, the District
            Attorney punished Zoe for exposing it.
          </p>
          <p>
            The District Attorney charged Zoe with a ramped-up felony conspiracy
            charge to increase her maximum sentence and forced her to wear a GPS
            ankle monitor for over 20 months while awaiting trial. When trial
            finally arrived, the court severely restricted Zoe&apos;s ability to
            present evidence of the animal cruelty she had documented at Perdue
            and reported to authorities before the rescue. The jury only heard a
            small sliver of the story and Zoe was ultimately convicted and taken
            immediately into custody.
          </p>
          <p>
            <b>
              Her incarceration isn&apos;t just unjust; it is life-threatening.
            </b>{" "}
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
              Compassion should not be criminalized, and it certainly
              shouldn&apos;t be a death sentence.
            </b>
          </p>
        </div>
        <div className="w-full lg:w-1/2 flex items-center justify-center order-1 lg:order-2">
          <Image
            src="https://worldanimalnews.com/wp-content/uploads/2025/09/IMG_5130.jpeg"
            alt="Exposing Perdue factory farm conditions"
            className="rounded-2xl w-full object-cover"
            style={{
              maxHeight: "500px",
            }}
            width={800}
            height={600}
          />
        </div>
      </div>
    </Section>
  );
}

function WhoIsZoeSection() {
  return (
    <Section>
      <h2 className="border-b border-slate-300 pb-2 uppercase text-xl tracking-wide text-slate-800">
        Who is Zoe Rosenberg?
      </h2>
      <div className="flex flex-col lg:flex-row items-center justify-evenly gap-12">
        <div className="w-full lg:w-1/2 flex items-center justify-center order-1 lg:order-1">
          <Image
            src={zoeWithAnimal}
            alt="Zoe with two chickens"
            className="rounded-2xl w-full object-cover"
            style={{ maxHeight: "500px" }}
          />
        </div>
        <div className="w-full lg:w-1/2 text-left space-y-8 order-2 lg:order-2">
          <p>
            Zoe is a 23-year-old college graduate of UC Berkeley, an animal
            sanctuary founder, and an animal cruelty investigator with Direct
            Action Everywhere. She is an incredible human being who has spent
            most of her young life working to protect animals.
          </p>
          <p>
            When she was just 11, Zoe founded Happy Hen Animal Sanctuary, a
            nonprofit that has saved more than 1,000 neglected or abandoned
            animals. Zoe has worked with law enforcement to rescue animals from
            neglect and abuse.
          </p>
          <p>
            She has bravely documented cruelty inside factory farms and rescued
            animals who would otherwise have died. And even throughout this
            unjust prosecution, she has continued to speak up boldly for
            animals.
          </p>
        </div>
      </div>
    </Section>
  );
}

function ZoesHealthSection() {
  return (
    <Section className="bg-red-50">
      <h2 className="border-b border-slate-300 pb-2 uppercase text-xl tracking-wide text-slate-800">
        Zoe&apos;s Health is at Immediate Risk
      </h2>
      <div className="flex flex-col lg:flex-row items-center justify-evenly gap-12">
        <div className="w-full lg:w-1/2 text-left space-y-8 order-2 lg:order-1">
          <p>
            Zoe lives with Type 1 Diabetes, meaning she is insulin dependent. If
            she goes off insulin for even just several hours, she could go into
            diabetic ketoacidosis which is life-threatening and requires
            immediate hospitalization.{" "}
            <b>
              Despite this, the Sonoma County jail will not allow Zoe to keep
              her insulin pump.
            </b>
          </p>
          <p>
            If her blood sugar drops without prompt treatment with glucose, she
            could have a seizure, go into a coma, or die. Zoe also suffers from
            serious complications from her diabetes, including gastroparesis –
            partial stomach paralysis that prevents her from eating or drinking
            enough, leading to dehydration, weight loss, and malnutrition.
          </p>
          <p>
            <b>
              Without constant access to her medical supplies and regular care,
              Zoe could die in custody.
            </b>{" "}
            Her incarceration is not only unjust—it is a direct threat to her
            life.
          </p>
          <p className="text-slate-700">
            Tragically, it is not uncommon for people with diabetes to die
            preventable deaths in custody due to lack of medical care.
          </p>
        </div>
        <div className="w-full lg:w-1/2 flex items-center justify-center order-1 lg:order-2">
          <Image
            src={zoeHealthRisk}
            alt="Zoe Rosenberg"
            className="rounded-2xl w-full object-cover max-h-80 lg:max-h-[500px]"
            style={{
              objectPosition: "center 40%",
            }}
          />
        </div>
      </div>
    </Section>
  );
}

function WhyIsZoeInJailSection() {
  return (
    <Section>
      <h2 className="border-b border-slate-300 pb-2 uppercase text-xl tracking-wide text-slate-800">
        Why is Zoe in Jail?
      </h2>
      <div className="flex flex-col lg:flex-row items-center justify-evenly gap-12">
        <div className="w-full lg:w-1/2 flex items-center justify-center order-1 lg:order-1">
          <Image
            src={zoeRescuedChicken}
            alt="Rescued chicken outdoors"
            className="rounded-2xl w-full object-cover max-h-80 lg:max-h-[700px] object-[center_60%] lg:object-[center_100%]"
          />
        </div>
        <div className="w-full lg:w-1/2 text-left space-y-8 order-2 lg:order-2">
          <p>
            In June 2023, Zoe and other animal cruelty investigators with Direct
            Action Everywhere (DxE) openly rescued four ailing chickens from
            Petaluma Poultry, a Perdue-owned slaughterhouse in Petaluma, CA.
          </p>
          <p>
            Those birds—<b>Poppy, Aster, Ivy, and Azalea</b>—were rushed to
            veterinary care. Poppy was treated with antibiotics for a
            respiratory infection, and Aster&apos;s infected feet were cleaned,
            treated, and bandaged. All four birds got the care they needed to
            recover, and today, they are thriving at an animal sanctuary.
          </p>
          <p>
            Before this rescue, Zoe and others had repeatedly reported unlawful
            animal cruelty at Petaluma Poultry, but authorities failed to act.
            Going all the way back to 2018, DxE investigations have documented
            routine violations of California&apos;s animal cruelty laws at
            Petaluma Poultry, including:
          </p>
          <ul className="list-disc pl-8 space-y-2">
            <li>
              Sick and injured birds collapsed on the factory farm floor, unable
              to reach food or water
            </li>
            <li>
              Open sores, wing and joint injuries, and necrotic wounds exposing
              muscle and bone
            </li>
            <li>
              Evidence of animals being scalded alive inside the slaughterhouse
            </li>
          </ul>
          <p>
            Petaluma Poultry, a subsidiary of national poultry giant Perdue
            Foods, supplies to major grocery chains like Safeway and Trader
            Joe&apos;s. The company&apos;s money and power has shielded it from
            facing consequences for clear crimes against animals, even when
            Sonoma County&apos;s Animal Services Department concurred with DxE
            and encouraged the Sheriff&apos;s Office to prosecute a Petaluma
            Poultry factory farm for animal cruelty. No prosecution ever
            happened.
          </p>
          <p>
            <b>
              Zoe took action to protect animals from cruelty when the
              authorities failed to do so.
            </b>{" "}
            She does not regret her actions. She knows that every suffering
            animal deserves to be rescued, and she will continue fighting for
            them.
          </p>
        </div>
      </div>
    </Section>
  );
}

function TakeActionSection() {
  return (
    <Section className="bg-slate-100">
      <h2 className="border-b border-slate-300 pb-2 uppercase text-xl tracking-wide text-slate-800 text-center">
        Take Action
      </h2>
      <div className="w-full text-left space-y-8">
        <p className="text-center text-2xl font-bold">
          Will you fight with her?
        </p>
        <ul className="list-disc pl-8 space-y-3 text-lg">
          <li>
            Ask your friends and family to sign the petition to free Zoe now
          </li>
          <li>
            Share Zoe&apos;s story on social media with{" "}
            <span className="font-semibold">#FreeZoe</span> and{" "}
            <span className="font-semibold">#RightToRescue</span>
          </li>
          <li>
            Join the animal rights movement. By locking Zoe up, prosecutors and
            industry are hoping to scare us into silence. Don&apos;t let them.
            Let&apos;s speak out louder than ever before.
          </li>
        </ul>
        <p className="text-center text-xl font-bold pt-4 pb-0">
          Every day Zoe spends behind bars is another day that her life is in
          jeopardy.
        </p>
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
