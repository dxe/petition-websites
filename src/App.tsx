import { HTMLProps, ReactNode, useCallback } from "react";
import heroDuckImg from "./assets/hero.jpeg";
import aboutImg1 from "./assets/about1.webp";
import aboutImg2 from "./assets/about2.webp";
import aboutImg3 from "./assets/about3.webp";
import { cn } from "./utils";
import { Petition } from "./components/petition.tsx";
import { Button } from "./components/ui/button.tsx";
import video1Thumb from "./assets/video1.jpeg";
import video2Thumb from "./assets/video2.jpeg";
import { PlayIcon } from "./components/svg/play-icon.tsx";
import {
  Dialog,
  DialogContent,
  DialogTrigger,
} from "./components/ui/dialog.tsx";

const App = () => {
  // TODO: add GA

  const scrollToPetition = useCallback(() => {
    const petitionSection = document.getElementById("petition-section");
    petitionSection?.scrollIntoView({ behavior: "smooth" });
  }, []);

  return (
    <div className="flex flex-col gap-6 items-center">
      <section
        className="md:min-h-[90vh] w-full text-white lg:bg-center md:bg-[40%] bg-[45%] bg-cover flex flex-col"
        style={{
          backgroundImage: `url(${heroDuckImg})`,
        }}
      >
        <div className="bg-black w-full flex-grow bg-opacity-40 flex justify-center items-center md:px-12 py-12">
          <div className="flex flex-col gap-9 max-w-screen-xl md:border-l-2 p-6 w-full">
            <div className="flex flex-col gap-4">
              <h1 className="uppercase text-2xl md:text-5xl leading-[1.125] tracking-wide max-w-[16rem] md:max-w-md border-b-2 md:border-0 py-4 md:py-0">
                The Reichardt Factory Farm Leaves Diseased Ducks to Die
              </h1>
              <p className="font-medium max-w-lg">
                Multiple investigations have exposed Reichardt Duck Farm for
                rampant disease and criminal animal cruelty.
              </p>
            </div>
            <Button
              className="self-start"
              variant="secondary"
              size="lg"
              onClick={scrollToPetition}
            >
              Tell the DA to prosecute Reichardt
            </Button>
          </div>
        </div>
      </section>
      <Section
        className="gap-12 items-center bg-slate-200 xl:rounded-lg py-12 md:px-16"
        id="petition-section"
      >
        <h2 className="font-semibold text-xl uppercase self-start text-center md:text-left w-full">
          Contact the District Attorney Now
        </h2>
        <Petition />
      </Section>
      <Section>
        <h2 className="border-b border-slate-300 pb-2 uppercase text-lg tracking-wide text-slate-800">
          About Reichardt Duck Farm
        </h2>
        <div className="flex flex-col lg:flex-row items-center justify-evenly gap-12 text-center">
          <div className="flex flex-col gap-6 max-w-xs w-full">
            <img src={aboutImg1} className="rounded-full" alt="Duck 1" />
            <div className="flex flex-col gap-4">
              <div className="text-4xl uppercase font-semibold">Disease</div>
              <p>
                Testing has found dangerous diseases are rampant at Reichardt,
                including E. coli, Riemerella anatipestifer, Salmonella,
                Staphylococcus, Aerococcus viridans, and Pseudomonas aeruginosa.
                Some of these bacteria could spread to humans. Reichardt has
                poor biosecurity practices which contribute to this rampant
                disease.
              </p>
            </div>
          </div>
          <div className="flex flex-col gap-6 max-w-xs w-full">
            <img src={aboutImg2} className="rounded-full" alt="Duck 2" />
            <div className="flex flex-col gap-4">
              <div className="text-4xl uppercase font-semibold">Neglect</div>
              <p>
                Undercover investigations have exposed that Reichardt leaves
                diseased ducks to die without veterinary care. Many ducks
                develop balance issues from infection and fall on their backs,
                unable to right themselves. Without help, they slowly die of
                dehydration and starvation.
              </p>
            </div>
          </div>
          <div className="flex flex-col gap-6 max-w-xs w-full">
            <img src={aboutImg3} className="rounded-full" alt="Duck 3" />
            <div className="flex flex-col gap-4">
              <div className="text-4xl uppercase font-semibold">Abuse</div>
              <p>
                Ducks at Reichardt are denied proper housing. They don't have
                access to water to swim or float in or for cleaning their eyes
                and sinuses. They are crowded together in filthy barns, forced
                to spend every moment living on a wire floor that digs into
                their feet and causes painful injuries. Ducklings routinely get
                stuck in the wire.
              </p>
            </div>
          </div>
        </div>
      </Section>
      <Section className="flex md:flex-row gap-8 text-center justify-evenly items-center bg-slate-200 xl:rounded-lg py-12 md:py-4">
        <div className="flex gap-6 flex-col items-center">
          <h4 className="text-3xl uppercase font-medium">
            Ducklings Trapped In Wire
          </h4>
          <p className="max-w-md">
            Investigators found dozens of ducklings trapped in wire on one night
            in October.
          </p>
          <Button
            size="lg"
            onClick={scrollToPetition}
            className="hidden md:flex"
          >
            Tell the DA to prosecute Reichardt
          </Button>
        </div>
        <Dialog>
          <DialogTrigger asChild>
            <button className="md:order-first max-w-[325px] relative">
              <PlayIcon className="inset-0 absolute w-full h-full scale-[33%]" />
              <img
                src={video1Thumb}
                className="aspect-square rounded-md"
                alt="Watch"
              />
            </button>
          </DialogTrigger>
          <DialogContent className="w-full flex justify-center bg-black">
            <iframe
              src="https://player.vimeo.com/video/899045165?h=2603019680&autoplay=1&title=0&byline=0&portrait=0&badge=0"
              allow="autoplay; fullscreen"
              allowFullScreen
              className="border-0 w-full h-full aspect-[9/16] md:order-first max-h-[80vh]"
            ></iframe>
          </DialogContent>
        </Dialog>
        <Button size="lg" onClick={scrollToPetition} className="flex md:hidden">
          Tell the DA to prosecute Reichardt
        </Button>
      </Section>
      <Section className="flex md:flex-row gap-8 text-center justify-evenly items-center bg-slate-200 xl:rounded-lg py-12 md:py-4">
        <div className="flex gap-6 flex-col items-center">
          <h4 className="text-3xl uppercase font-medium">Meet River</h4>
          <p className="max-w-md">
            River was on the verge of death at Reichardt Duck Farm.
            Investigators knew they couldn't leave him behind so they rushed him
            to the vet.
          </p>
          <Button
            size="lg"
            onClick={scrollToPetition}
            className="hidden md:flex"
          >
            Tell the DA to prosecute Reichardt
          </Button>
        </div>
        <Dialog>
          <DialogTrigger asChild>
            <button className="md:order-first max-w-[325px] relative">
              <PlayIcon className="inset-0 absolute w-full h-full scale-[33%]" />
              <img
                src={video2Thumb}
                className="aspect-square rounded-md"
                alt="Watch"
              />
            </button>
          </DialogTrigger>
          <DialogContent className="w-full flex justify-center bg-black">
            <iframe
              src="https://player.vimeo.com/video/899042025?h=d2d319d36b&autoplay=1&title=0&byline=0&portrait=0&badge=0"
              allow="autoplay; fullscreen"
              allowFullScreen
              className="border-0 w-full h-full aspect-[9/16] md:order-first"
            ></iframe>
          </DialogContent>
        </Dialog>
        <Button size="lg" onClick={scrollToPetition} className="flex md:hidden">
          Tell the DA to prosecute Reichardt
        </Button>
      </Section>
      <Section className="text-center pt-4 pb-12 text-sm md:border-t border-slate-300 text-slate-600">
        &copy; {new Date().getFullYear()} Help The Ducks
      </Section>
    </div>
  );
};

export default App;

const Section = ({
  children,
  className,
  ...props
}: {
  children: ReactNode;
  className?: string;
} & HTMLProps<HTMLDivElement>) => {
  return (
    <section
      {...props}
      className={cn(
        "flex flex-col gap-8 max-w-screen-xl w-full p-4",
        className,
      )}
    >
      {children}
    </section>
  );
};
