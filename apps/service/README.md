# Email petition service

For helptheducks.com, helpthechickens.com, and righttorescue.com.

Run the dev server:
`go run .`

## Configuration

The service uses environment variables for configuration. There are two types of environment files:

### .env.shared
Contains shared environment variables that are checked into version control:
- `PORT`: Server port (default: 3333)
- `DB_CONNECTION_STRING`: Database connection string
- `RECAPTCHA_SECRET`: reCAPTCHA secret key
- `GO_ENV`: Environment (development/production)

## Build Process

During build time (strictly in development environment), the service merges environment variables from both `.env.shared` and `.env.local` files using a tasks.json script. This ensures that:

1. Shared configuration is reproducible to all developers
2. Sensitive API keys remain local and secure
3. The testing deployment has all necessary environment variables