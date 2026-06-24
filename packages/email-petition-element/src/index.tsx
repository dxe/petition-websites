import styles from "./styles.css?inline";
import r2wc from "@r2wc/react-to-web-component";
import { EmailPetition } from "@dxe/email-petition/email-petition";

/**
 * Tailwind v4 implements many utilities (borders, rings, shadows, transforms,
 * gradients, …) via `@property`-registered custom properties such as
 * `--tw-border-style`. Chromium ignores `@property` rules that live inside a
 * shadow root, so those variables have no registered initial value and the
 * dependent utilities silently break — e.g. `.border` resolves to
 * `border-style: var(--tw-border-style)` → invalid → `none`, so input borders
 * disappear.
 *
 * `@property` rules only declare a custom property's type/default; they apply no
 * visual styling. Registering them once at the document level therefore fixes
 * the cascade inside the shadow root without leaking any appearance to the host
 * page (which never references these `--tw-*` properties).
 */
function registerTailwindProperties(css: string) {
  const id = "email-petition-tw-property-shim";
  if (document.getElementById(id)) return;
  const propertyRules = css.match(/@property[^{]+\{[^}]*\}/g);
  if (!propertyRules) return;
  const style = document.createElement("style");
  style.id = id;
  style.textContent = propertyRules.join("");
  document.head.appendChild(style);
}
registerTailwindProperties(styles);

// Thin wrapper that flattens signatureThermometer into a simple numeric attribute.
function EmailPetitionWrapper(props: {
  petitionId: string;
  campaignName: string;
  defaultMessage: string;
  debug?: boolean;
  test?: boolean;
  thermometerGoal?: number;
}) {
  const required = {
    "petition-id": props.petitionId,
    "campaign-name": props.campaignName,
    "default-message": props.defaultMessage,
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
      <style>{styles}</style>
      <EmailPetition
        petitionId={props.petitionId ?? ""}
        campaignName={props.campaignName ?? ""}
        defaultMessage={props.defaultMessage ?? ""}
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
    debug: "boolean",
    test: "boolean",
    thermometerGoal: "number",
  },
});

if (!customElements.get("email-petition")) {
  customElements.define("email-petition", EmailPetitionElement);
}
