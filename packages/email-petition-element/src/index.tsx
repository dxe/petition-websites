import "./styles.css";
import r2wc from "@r2wc/react-to-web-component";
import { EmailPetition } from "@dxe/email-petition/email-petition";

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
  );
}

const EmailPetitionElement = r2wc(EmailPetitionWrapper, {
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
