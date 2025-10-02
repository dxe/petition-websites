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

### Image optimization

Convert any new PNGs to JPGs. Make sure JPGs are a reasonable size.

### Mail petition service

Update the service with configuration to send mail from the petition domain.

### Captcha

Add the petition domain to the captcha in Google reCAPTCHA Admin.

### Create new identity in AWS SES

Go to AWS SES, go to the identities section and create a new "domain"-type idenity and set the MAIL FROM to
`mail.<petition-domain>`.
Verify using Easy DKIM and choose RSA_2048_BIT key length.

Check Namecheap to make sure the nameservers point to AWS so we can use Route 53 for DNS.
SES will create the necessary records in Route 53 to verify the domain by default.

Make sure SES finishes verifying the domain.

This is just for identifying the server that sends mail.
No mail is actually sent from or received at this domain.

### Create new bucket in AWS S3 (for dedicated petition website only)

Allow all public access unless you want to figure out permissions for deployment and CloudFront.

Under properties, enable static website hosting.
Set the index document to `index.html`.

### Create new distribution in CloudFront (for dedicated petition website only)

Target the S3 bucket's website endpoint. If turned on static website hosting
in S3, AWS will prompt you to set it correctly when you update the "origin
domain" field.

Example:

- Incorrect:
  - `helpthechickens.com.s3.us-west-2.amazonaws.com`
- Correct:
  - `helpthechickens.com.s3-website-us-west-2.amazonaws.com`

Set the alternate domain name to the petition domain (without any www).

For "Custom SSL certificate" certificate, click "Request certificate", and
create a certificate for www. and @ domains, and verify ownership with DNS
records.

Set the price class to "Use only North America and Europe" (do not use all edge
locations).

Set behavior -> viewer protocol policy: "Redirect HTTP to HTTPS".

Check the box to enable IPv6 support.

All other settings left as default.

### Point domain to CloudFront (for dedicated petition website only)

Go to namecheap advanced dns for the domain, create CNAME record `@` pointing
to the domain name of the CloudFront distribution found in AWS Console on the
"General" tab for the distribution.

### Create and run GitHub workflow (for dedicated petition website only)

Create a workflow file in .github/workflows and follow the other files as an
example.

Add variables and secrets in the GitHub repo settings:

- https://github.com/dxe/helptheducks.com/settings/variables/actions
- https://github.com/dxe/helptheducks.com/settings/secrets/actions

Run the workflow by pushing to the main branch or running it manually from here:

- https://github.com/dxe/helptheducks.com/actions

### TODO (for dedicated petition website only)

- forwarding `www.` domain with S3/CloudFront or Cloudflare
