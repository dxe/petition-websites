# DxE Petitions Sites monorepo

Monolithic (multi-project / multi-package) repository for petition websites.

Note: some petition websites may still have their own repos.

## monorepo setup

We use Turborepo to manage this monorepo of petition websites.

Note Turbo currently ignores the Go project in apps/service since there is no
package.json.

### Build

First install dependencies:

```
pnpm install
```

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

## Environment variables

Each petition website has `.env.development` and `.env.production` environment
variable files. The development file is used with `pnpm dev` and production with
`pnpm build && pnpm start`.

Below are variables used to configure the websites.

`NEXT_PUBLIC_` prefix allows these variables to be substituted during build time
into the static files served publicly to users' browsers.

### `NEXT_PUBLIC_PETITIONS_API_ROOT`

Base URL for the general petition service. This records tallys and then invokes
the signup service to subscribe the user to emails.

### `NEXT_PUBLIC_PETITION_ID`

The petition ID used by the petition service.

Prefix with `test:` to avoid contaminating the tally for the
real petition. The prefix is removed

### `NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT`

Base URL for the petition mail service. It sends the emails to our targets.
It also records a tally of its own and records copies of messages sent.

### `NEXT_PUBLIC_CAMPAIGN_NAME`

The name of the campaign used by the petition mail service. Configured in
`apps/service/config/config.go`.

Points to configuration for emails such as sender and target emails, and subject
line. For testing, use the value `test`.

## Creating a new petition site

Start early on the AWS setup in case of any delays for domain verification for
SES (email) or ACM (cert for CloudFront).

Note in AWS the region we use is `us-west-2`.

### Buy new domain (for dedicated petition website only)

Buy a domain for the petition in Namecheap.

### Copy existing site (for dedicated petition website only)

1. Make a copy of an existing petition website, such as the most recently created one.
2. Choose an unused port number in package.json in the directory for the new website.
3. Update the petition component props, Google Analytics tag, `.env` files,
   layout file and homepage files, images and any other references to the old
   petition.

See [environment variables](#environment-variables) for details about the `.env`
files.

### Update launch.json

Update .vscode/launch.json file for easier dubugging in VS Code.

### Image optimization

Convert any new PNGs to JPGs. Make sure JPGs are a reasonable size.

### Mail petition service

Update the service with configuration to send mail from the petition domain. Look for `AllowedOrigins` in `main.go` and
`EmailSettings` in `config/config.go`.

### Captcha

Add the petition domain to the captcha in Google reCAPTCHA Admin at https://www.google.com/recaptcha/admin/

### Point domain nameservers to Route 53

Check Namecheap to make sure the nameservers point to AWS so we can use Route 53 for DNS.

Go to Route 53 and create a "hosted zone" for the petition domain. Put the provided NS records into Namecheap.

### Create new identity in AWS SES

Go to AWS SES, go to the identities section and create a new "domain"-type idenity and set the MAIL FROM to
`mail.<petition-domain>`.
Verify using Easy DKIM and choose RSA_2048_BIT key length.

After creation, under "DomainKeys Identified Mail (DKIM)" section, click the "Publish records to Route 53" button.
SES will create the necessary records in Route 53 to verify the domain by default.

Make sure SES finishes verifying the domain.

This is just for identifying the server that sends mail.
No mail is actually sent from or received at this domain.

### Create new bucket in AWS S3 (for dedicated petition website only)

Create a bucket in S3 named after the petition domain

Leave all the default settings. It's ok to leave "Block all public access" checked, and do NOT enable static website
hosting as it will be hosted via CloudFront instead.

### Create new distribution in CloudFront (for dedicated petition website only)

Set the domain name to the petition domain (without www), verify, click "Check" and add a subdomain for "www".

Target the S3 bucket. Follow its instructions to create and attach an Origin Access Control, and copy the policy it
gives you to the S3 bucket's permissions page.

No need for Web Application Firewall (WAF) at this time.

In the "Get TLS certificate" section, check both boxes to include the apex domain and all subdomains (wildcard) and
click "Create Certificate".

After creation:

Make sure both www.@ and @ are listed under "alternative domain names", then click "Route domains to CloudFront".

Edit the "Default root object" to be `index.html`.

Edit the preexisting "behavior":

- Ensure it has viewer protocol policy: "Redirect HTTP to HTTPS".
- Set "viewer request" to "CloudFront Functions" -> "redirect-to-www" or "redirect-from-www".

All other settings left as default.

No longer applicable, or maybe only applicable if not using Route 53?

- Set the price class to "Use only North America and Europe" (do not use all edge
  locations).
- Check the box to enable IPv6 support.

### Create and run GitHub workflow (for dedicated petition website only)

Create a workflow file in .github/workflows and follow the other files as an
example.

Add variables and secrets in the GitHub repo settings:

- https://github.com/dxe/petition-websites/settings/variables/actions
- https://github.com/dxe/petition-websites/settings/secrets/actions

You'll need to add variables that follow the pattern below and use them in the workflow file:

- `*_CLOUDFRONT_DISTRIBUTION`
- `*_S3_BUCKET`

Run the workflow by pushing to the main branch or running it manually from here:

- https://github.com/dxe/helptheducks.com/actions

### TODO (for dedicated petition website only)

- forwarding `www.` domain with S3/CloudFront or Cloudflare
