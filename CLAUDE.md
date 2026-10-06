# CLAUDE.md

## Websites

The `apps` directory contains campaign websites. They are statically generated Next.js sites.

## Email service

The Go backend that sends, tallies and records the email sent is in `apps/service/`.

Note our petitions post to both this service as well as another service outside this repo called the "petition service" which records its own tally and handles SendGrid email list integration.

To change recipient addresses and subject lines of email petitions, see `apps/service/config/config.go`.

## Components

The reusable petition component (React) is in `packages/email-petition`. It is wrapped as a custom element in `packages/email-petition-element`.

Lower-level components are in `packages/components`.

## Formatting

When editing or creating files in TypeScript projects, run the Prettier auto-formatter.

Format one file: `pnpm exec prettier --write <files>`

Format all files: `pnpm format:fix`
