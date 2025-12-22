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

### .env.local
Contains local environment variables that should NOT be checked into version control:
- `GOOGLE_MAPS_GEOCODING_API_KEY`: Google Maps Geocoding API key for ZIP code to city lookup

**IMPORTANT**: `GOOGLE_MAPS_GEOCODING_API_KEY` must be placed in `.env.local` and should never be committed to version control.

## Build Process

During build time, the service merges environment variables from both `.env.shared` and `.env.local` files using a tasks.json script. This ensures that:

1. Shared configuration is available to all developers
2. Sensitive API keys remain local and secure
3. The final deployment has all necessary environment variables

### Build Script
The build process uses a script that combines:
```bash
cat .env.shared .env.local > .env.debug
```

This creates a temporary `.env.debug` file containing all environment variables for the build process.

## API Keys Required

### For City Autocomplete Mode

When using `citySelectionMode="autocompleteTextbox"` in the frontend, the following API keys are required:

1. **Frontend (in website .env files):**
   - `NEXT_PUBLIC_GOOGLE_MAPS_PLACES_NEW_API_KEY`: For Google Places API (New) autocomplete
   - Location: `apps/[website]/.env.development` or `apps/[website]/.env.local`

2. **Backend (in service .env.local):**
   - `GOOGLE_MAPS_GEOCODING_API_KEY`: For Google Geocoding API ZIP lookup
   - Location: `apps/service/.env.local`

### Setup Instructions

1. Create `apps/service/.env.local` with your Google Maps Geocoding API key:
   ```
   GOOGLE_MAPS_GEOCODING_API_KEY=your_geocoding_api_key_here
   ```

2. Ensure frontend apps have their respective API keys in their environment files

3. The build process will automatically merge the environment files during deployment

## Security Notes

- Never commit `.env.local` files to version control
- API keys should be kept secure and only shared with authorized team members
- Use different API keys for development and production environments when possible
