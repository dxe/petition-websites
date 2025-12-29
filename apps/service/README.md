# Email petition service

For helptheducks.com, helpthechickens.com, freezoe.org, and righttorescue.com.

Run the dev server:
`go run .`

## Configuration

The service uses environment variables for configuration. Environment variables are loaded from two files when using VS Code debugger: .env.shared (version controlled) and .env.local (not version controlled, better for secrets).

The following variables are expected in .env.local:
 - `GOOGLE_MAPS_GEOCODING_API_KEY`
 
 See `launch.json` for details on the mechanics of loading the `.env` files.