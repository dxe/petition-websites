# @dxe/email-petition

The React component (`EmailPetition`) that renders the petition sign / send-a-message
form. It collects the signer's details, runs an invisible reCAPTCHA, and submits to
the two backend services (the petition tally API and the campaign mailer API).

This package is the source of truth for the form's behavior. It is consumed two ways:

- directly by the Next.js petition websites in `apps/*`, and
- wrapped as a framework-agnostic `<email-petition>` web component by
  [`@dxe/email-petition-element`](../email-petition-element/README.md) for embedding
  on third-party sites.

## Usage

```tsx
import { EmailPetition } from "@dxe/email-petition/email-petition";

<EmailPetition
  petitionId="your-petition-id"
  campaignName="your-campaign-name"
  defaultMessage={
    "Dear Decision Maker,\n\nI urge you to take action.\n\n[Your name], [Your city]"
  }
  locationInputMode="zipWithSonomaCountyCity"
  debug={false}
  test={false}
/>;
```

### Props

| Prop                   | Type                      | Required | Description                                                                                  |
| ---------------------- | ------------------------- | -------- | -------------------------------------------------------------------------------------------- |
| `petitionId`           | `string`                  | yes      | ID sent to the petition API as `id`.                                                         |
| `campaignName`         | `string`                  | yes      | Campaign name sent to the mailer API as `campaign`.                                          |
| `defaultMessage`       | `string`                  | yes      | Pre-filled message body. Supports `[Your name]` and `[Your city]` placeholders (see below).  |
| `locationInputMode`    | `LocationInputMode`       | yes      | Which location fields to collect. See [Location input modes](#location-input-modes).         |
| `debug`                | `boolean`                 | yes      | Logs the resolved API URLs and IDs to the console on render.                                 |
| `test`                 | `boolean`                 | yes      | Prefixes `petitionId` and `campaignName` with `test:` so submissions don't affect prod data. |
| `signatureThermometer` | `{ defaultGoal: number }` | no       | Renders a signature-count thermometer with the given goal.                                   |
| `onSubmit`             | `() => void`              | no       | Called when the user submits (before the network requests).                                  |

`LocationInputMode` is exported from this package for typing.

### Message placeholders

`[Your name]` and `[Your city]` in `defaultMessage` are substituted with the signer's
entered values. The substitution stops once the user edits the message themselves, so
their customizations are preserved.

### Location input modes

The `locationInputMode` prop controls how a signer's location is collected and what
gets sent to the APIs:

- `zipWithSonomaCountyCity`: a US zip code field, an auto-derived city selector for
  Sonoma County zips, and an "Outside the United States" toggle.
- `sfOnly`: hides zip and city and instead shows a required "I am a resident of San
  Francisco, CA" checkbox. Submissions are tagged with San Francisco, CA, US.

## Backend requests

On submit, the component obtains a reCAPTCHA token and then posts to the two services
**sequentially** — the petition API first, then the mailer. If the petition request
fails, the mailer request is never sent, so the user can safely resubmit without
sending duplicate emails.

The endpoint base URLs come from environment variables (see below):

| Service  | URL                                                      | Encoding                            |
| -------- | -------------------------------------------------------- | ----------------------------------- |
| Petition | `${NEXT_PUBLIC_PETITIONS_API_ROOT}/sign`                 | `application/x-www-form-urlencoded` |
| Mailer   | `${NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT}/message/create` | `application/json`                  |

**Petition (`/sign`)** fields: `id`, `name`, `email`, `phone` (if provided), the
resolved location fields (`zip` / `city` / `country`, or the fixed San Francisco
values in `sfOnly`), and `fullHref` (the current page URL).

**Mailer (`/message/create`)** fields: `name`, `email`, `phone` (if provided),
`outside_us`, the resolved location fields, the substituted `message`, `campaign`,
and the reCAPTCHA `token`.

### Environment variables

The component reads these at module load to build the endpoint URLs. They are
provided by the consuming app at build time (see the
[root README](../../README.md#environment-variables) for the full list and how the
apps configure them):

- `NEXT_PUBLIC_PETITIONS_API_ROOT`
- `NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT`

The reCAPTCHA site key is hardcoded in this package.

## Development

```sh
pnpm install
pnpm build        # one-off compile to dist/
pnpm dev
```

To see changes live in a website's dev server, run the website and this package
together with Turbo filters from the repo root — see the
[root README](../../README.md#develop).
