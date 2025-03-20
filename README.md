# DxE Petitions Sites monorepo

Monolithic (multi-project / multi-package) repository for petition websites.

Note: some petition websites may still have their own repos.

## monorepo setup

We use Turborepo to manage this monorepo of petition websites.

Note Turbo currently ignores the Go project in apps/service since there is no
package.json.

### Build

To build all apps and packages, run the following command:

```
pnpm build
```

### Develop

Run the dev server for one website:

```bash
pnpm dev --filter helptheducks.com`
```

Or: `cd apps/helptheducks.com` and `pnpm dev`.

Run the dev server for one website and some dependencies:

```bash
pnpm dev --filter helptheducks.com --filter @dxe/email-petition --filter @dxe/petitions-components
```

The above command will make changes to the email-petition component become live
immediately in the browser window for helptheducks website dev server.

Develop all apps and packages together:

```bash
pnpm dev
```

### Turborepo Links

Learn more about Turborepo:

- [Tasks](https://turbo.build/repo/docs/core-concepts/monorepos/running-tasks)
- [Caching](https://turbo.build/repo/docs/core-concepts/caching)
- [Remote Caching](https://turbo.build/repo/docs/core-concepts/remote-caching)
- [Filtering](https://turbo.build/repo/docs/core-concepts/monorepos/filtering)
- [Configuration Options](https://turbo.build/repo/docs/reference/configuration)
- [CLI Usage](https://turbo.build/repo/docs/reference/command-line-reference)
