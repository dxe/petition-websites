import styles from "./styles.css?inline";
import r2wc from "@r2wc/react-to-web-component";
import { EmailPetition } from "@dxe/email-petition/email-petition";
import type { LocationInputMode } from "@dxe/email-petition/email-petition";

/**
 * Tailwind v4 implements many utilities (borders, rings, shadows, transforms,
 * gradients, …) via `@property`-registered custom properties such as
 * `--tw-border-style`. Chromium ignores `@property` rules that live inside a
 * shadow root, so those variables lose their registered initial value and the
 * dependent utilities silently break — e.g. `.border` resolves to
 * `border-style: var(--tw-border-style)` → invalid → `none`, so input borders
 * disappear.
 *
 * Rather than registering the properties at the document level (which would
 * mutate a global, Tailwind-shared namespace and could affect host pages), we
 * re-create the lost defaults as ordinary custom properties scoped to the shadow
 * tree: read each `@property` initial-value and emit it on every element inside
 * `@layer base`. The utilities that set these vars live in `@layer utilities`
 * (higher priority than base), so per-element overrides like `border-dashed`
 * still win. This mirrors the fallback Tailwind itself ships for Firefox, and
 * leaves the host document completely untouched.
 */
function withShadowPropertyDefaults(css: string): string {
  const decls: string[] = [];
  const propertyRule = /@property\s+(--[\w-]+)\s*\{([^}]*)\}/g;
  for (
    let match = propertyRule.exec(css);
    match;
    match = propertyRule.exec(css)
  ) {
    const initialValue = match[2].match(/initial-value:\s*([^;]+)/);
    if (initialValue) decls.push(`${match[1]}:${initialValue[1].trim()}`);
  }
  if (decls.length === 0) return css;
  return `${css}\n@layer base{*,::after,::before,::backdrop,::file-selector-button{${decls.join(";")}}}`;
}

const shadowStyles = withShadowPropertyDefaults(styles);

// Thin wrapper that flattens signatureThermometer into a simple numeric attribute.
function EmailPetitionWrapper(props: {
  petitionId: string;
  campaignName: string;
  defaultMessage: string;
  locationInputMode: LocationInputMode;
  debug?: boolean;
  test?: boolean;
  thermometerGoal?: number;
}) {
  const required = {
    "petition-id": props.petitionId,
    "campaign-name": props.campaignName,
    "default-message": props.defaultMessage,
    "location-input-mode": props.locationInputMode,
  };
  for (const [attribute, value] of Object.entries(required)) {
    if (!value) {
      throw new Error(
        `<email-petition> is missing required attribute: ${attribute}`,
      );
    }
  }

  return (
    <>
      {/* Rendered into the shadow root, so the styles are scoped to this
          component and don't leak to (or inherit from) the host page. */}
      <style>{shadowStyles}</style>
      <EmailPetition
        petitionId={props.petitionId}
        campaignName={props.campaignName}
        defaultMessage={props.defaultMessage}
        locationInputMode={props.locationInputMode}
        debug={props.debug ?? false}
        test={props.test ?? false}
        signatureThermometer={
          props.thermometerGoal != null
            ? { defaultGoal: props.thermometerGoal }
            : undefined
        }
      />
    </>
  );
}

const EmailPetitionElement = r2wc(EmailPetitionWrapper, {
  shadow: "open",
  props: {
    petitionId: "string",
    campaignName: "string",
    defaultMessage: "string",
    locationInputMode: "string",
    debug: "boolean",
    test: "boolean",
    thermometerGoal: "number",
  },
});

if (!customElements.get("email-petition")) {
  customElements.define("email-petition", EmailPetitionElement);
}
