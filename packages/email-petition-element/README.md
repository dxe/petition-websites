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

To preview the component while developing, run a dev server that serves the test / example `index.html` page:

```sh
pnpm serve
```

This starts Vite (with `--host`) and serves the example `index.html`, which loads the component so you can see and interact with it in the browser.

In VS Code you can launch the same thing from the **Run and Debug** panel with the **"Dev email petition element test page"** configuration, which runs `pnpm run serve` in this package.

## Usage

```html
<script src="/path/to/email-petition.js"></script>

<email-petition
  petition-id="your-petition-id"
  campaign-name="your-campaign-name"
  default-message="Dear Decision Maker,&#10;&#10;I urge you to take action.&#10;&#10;[Your name], [Your city]"
></email-petition>
```

### With the signature thermometer

```html
<email-petition
  petition-id="your-petition-id"
  campaign-name="your-campaign-name"
  default-message="Dear Decision Maker, ..."
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
  test
  debug
></email-petition>
```

## Attribute reference

| Attribute          | Type    | Required | Description                                                                    |
| ------------------ | ------- | -------- | ------------------------------------------------------------------------------ |
| `petition-id`      | string  | yes      | ID passed to the petition signing API                                          |
| `campaign-name`    | string  | yes      | Campaign name passed to the mailer API                                         |
| `default-message`  | string  | yes      | Pre-filled message body; supports `[Your name]` and `[Your city]` placeholders |
| `thermometer-goal` | number  | no       | Shows a signature thermometer with this goal count                             |
| `test`             | boolean | no       | Prefixes IDs with `test:` to avoid polluting production data                   |
| `debug`            | boolean | no       | Logs resolved API URLs and IDs to the console                                  |

## Notes

- The component is rendered into the **light DOM** (no shadow DOM), so your page's global styles will apply alongside the bundled Tailwind CSS.
- API URLs are baked in at build time via the `.env` file — they are not runtime-configurable through attributes.
- The reCAPTCHA site key is hardcoded in the underlying `@dxe/email-petition` package.
