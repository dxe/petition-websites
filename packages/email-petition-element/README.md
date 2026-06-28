# email-petition-element

A self-contained web component (`<email-petition>`) that wraps the `@dxe/email-petition` React component. Drop a single `<script>` tag into any HTML page — no React, no build tooling required on the consumer side.

## Building

Install dependencies and build:

```sh
pnpm install
pnpm build
```

The output is a single file: `dist/email-petition.js`. CSS is injected automatically at runtime — no separate stylesheet needed.

## Local preview

There are two ways to preview the component, and they test different things.

### 1. Dev server (fast iteration)

```sh
pnpm serve
```

This starts Vite (with `--host`) and opens the example page from `dev-preview/`
(served at `/`), which loads the component's **source** (`src/index.tsx`) as
native ES modules with hot reloading. Use this for fast iteration on the
component itself.

In VS Code you can launch the same thing from the **Run and Debug** panel with the
**"Dev email petition element preview - fast reload"** configuration, which runs
`pnpm run serve` in this package.

The dev server loads the source as modules through Vite — it does **not** run the
bundled IIFE that ships to consumers. So it won't catch bugs that only exist in the
built artifact (undefined globals like `process`, bundling/`define` gaps, etc.). For
that, use the built-artifact preview below.

### 2. Built-artifact preview (emulates a third-party site)

```sh
pnpm preview
```

This builds the bundle and serves the `build-preview/` directory over a plain static
HTTP server (no Vite, no module layer). Open <http://localhost:4173/> — it loads the
built `email-petition.js` via a plain `<script src>`, exactly the way a third-party
site embeds it. This is the faithful way to catch bugs that only appear in the shipped
IIFE before they reach production.

In VS Code you can launch this from the **Run and Debug** panel with the
**"Dev email petition element preview - build artifact"** configuration.

> Requires `python3` (preinstalled in the dev container). Stop the server with `Ctrl+C`.

## Usage

```html
<script src="https://s3.dxe.io/email-petition/email-petition.js"></script>

<email-petition
  petition-id="your-petition-id"
  campaign-name="your-campaign-name"
  default-message="Dear Decision Maker,&#10;&#10;I urge you to take action.&#10;&#10;[Your name], [Your city]"
  location-input-mode="zipWithSonomaCountyCity"
></email-petition>
```

### With the signature thermometer

```html
<email-petition
  petition-id="your-petition-id"
  campaign-name="your-campaign-name"
  default-message="Dear Decision Maker, ..."
  location-input-mode="zipWithSonomaCountyCity"
  thermometer-goal="1000"
></email-petition>
```

### Test / debug mode

Use `test` to prefix IDs with `test:` so submissions don't affect production data. Use `debug` to log API URLs and IDs to the console.

```html
<email-petition
  petition-id="my-petition"
  campaign-name="my-campaign"
  default-message="..."
  location-input-mode="zipWithSonomaCountyCity"
  test
  debug
></email-petition>
```

## Attribute reference

| Attribute             | Type    | Required | Description                                                                       |
| --------------------- | ------- | -------- | --------------------------------------------------------------------------------- |
| `petition-id`         | string  | yes      | ID passed to the petition signing API                                             |
| `campaign-name`       | string  | yes      | Campaign name passed to the mailer API                                            |
| `default-message`     | string  | yes      | Pre-filled message body; supports `[Your name]` and `[Your city]` placeholders    |
| `location-input-mode` | string  | yes      | Which location fields to collect. `zipWithSonomaCountyCity` or `sfOnly`            |
| `thermometer-goal`    | number  | no       | Shows a signature thermometer with this goal count                                |
| `test`                | boolean | no       | Prefixes IDs with `test:` to avoid polluting production data                      |
| `debug`               | boolean | no       | Logs resolved API URLs and IDs to the console                                     |

### Location input modes

The `location-input-mode` attribute is required and controls how a signer's
location is collected:

- `zipWithSonomaCountyCity`: a US zip code field, an auto-derived city
  selector for Sonoma County zips, and an "Outside the United States" toggle.
- `sfOnly`: hides zip and city and instead shows a required "I am a resident of
  San Francisco, CA" checkbox. Submissions are tagged with San Francisco, CA, US.

```html
<email-petition
  petition-id="your-petition-id"
  campaign-name="your-campaign-name"
  default-message="..."
  location-input-mode="sfOnly"
></email-petition>
```

## Notes

- The component is rendered into a **shadow DOM** root, so page-global styles do not automatically leak in; styles are bundled with the element.
- API URLs are baked in at build time via the `.env` file — they are not runtime-configurable through attributes.
- The reCAPTCHA site key is hardcoded in the underlying `@dxe/email-petition` package.
