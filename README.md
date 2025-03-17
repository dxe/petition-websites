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

To develop all apps and packages, run the following command:

```
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

## Creating a new petition site

Start early on the AWS setup in case of any delays for domain verification for
SES (email) or ACM (cert for CloudFront).

Note in AWS the region we use is `us-west-2`.

### Buy new domain

Buy a domain for the petition in Namecheap.

### Copy existing site

Make a copy of an existing petition website, such as the most recently created
one. Update the petition component props, Google Analytics tag, `.env` files,
layout file and homepage files, images and any other references to the old
petition.

### Image optimization

Convert any new PNGs to JPGs. Make sure JPGs are a reasonable size.

### Mail petition service

Update the service with configuration to send mail from the petition domain.

### Captcha

Add the petition domain to the captcha in Google reCAPTCHA Admin.

### Create new identity in AWS SES

Set up MX records.

Set up and verify DKIM.

Create and verify a custom MAIL FROM with domain `mail.<petition-domain>`.
This is just for identifying the server that sends mail. No mail is actually
sent from or received at this domain.

### Create new bucket in AWS S3

Allow all public access unless you want to figure out permissions for deployment
and CloudFront.

Under properties, enable static website hosting.
Set the index document to `index.html`.

### Create new distribution in CloudFront

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

### Point domain to CloudFront

Go to namecheap advanced dns for the domain, create CNAME record `@` pointing
to the domain name of the CloudFront distribution found in AWS Console on the
"General" tab for the distribution.

### Create and run GitHub workflow

Create a workflow file in .github/workflows and follow the other files as an
example.

Add variables and secrets in the GitHub repo settings:

- https://github.com/dxe/helptheducks.com/settings/variables/actions
- https://github.com/dxe/helptheducks.com/settings/secrets/actions

Run the workflow by pushing to the main branch or running it manually from here:

- https://github.com/dxe/helptheducks.com/actions

### TODO

- forwarding `www.` domain with Cloudflare.
