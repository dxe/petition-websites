# Google Maps Geocoding API Documentation

## Overview
The Google Maps Geocoding API converts addresses (including ZIP codes) into geographic coordinates and structured address components. This service is used in the petition system to validate ZIP codes and extract city information with optional geographic scope validation.

## API Endpoint
```
GET https://maps.googleapis.com/maps/api/geocode/json
```

## Request Parameters

### Basic ZIP Code Lookup
```
https://maps.googleapis.com/maps/api/geocode/json?address=95401&components=country:US&key=YOUR_API_KEY
```

### Required Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `address` | string | Yes | Address or ZIP code to geocode |
| `key` | string | Yes | Valid Google Maps API key |

### Optional Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `components` | string | Component filtering (e.g., `country:US`) |
| `language` | string | Response language (default: English) |
| `region` | string | Region biasing (e.g., `us`) |

## Response Structure

### Successful Response
```json
{
  "results": [
    {
      "address_components": [
        {
          "long_name": "Santa Rosa",
          "short_name": "Santa Rosa", 
          "types": ["locality", "political"]
        },
        {
          "long_name": "Sonoma County",
          "short_name": "Sonoma County",
          "types": ["administrative_area_level_2", "political"]
        },
        {
          "long_name": "California",
          "short_name": "CA",
          "types": ["administrative_area_level_1", "political"]
        },
        {
          "long_name": "United States",
          "short_name": "US", 
          "types": ["country", "political"]
        },
        {
          "long_name": "95401",
          "short_name": "95401",
          "types": ["postal_code"]
        }
      ],
      "formatted_address": "Santa Rosa, CA 95401, USA",
      "geometry": {
        "location": {
          "lat": 38.4405,
          "lng": -122.7144
        },
        "location_type": "APPROXIMATE",
        "viewport": {
          "northeast": {
            "lat": 38.5429,
            "lng": -122.6130
          },
          "southwest": {
            "lat": 38.3381,
            "lng": -122.8158
          }
        }
      },
      "place_id": "ChIJr8p0nB2mj4ARP5yVCLLkQzA",
      "types": ["postal_code"]
    }
  ],
  "status": "OK"
}
```

### Error Response
```json
{
  "results": [],
  "status": "ZERO_RESULTS"
}
```

## Address Component Types

The API returns structured address components with the following types relevant to the petition system:

| Type | Description | Example |
|------|-------------|---------|
| `locality` | City/town name | "Santa Rosa" |
| `administrative_area_level_1` | State/province | "California" |
| `administrative_area_level_2` | County/region | "Sonoma County" |
| `country` | Country name | "United States" |
| `postal_code` | ZIP/postal code | "95401" |

## Status Codes

| Status | Description | Action |
|--------|-------------|--------|
| `OK` | Request successful | Process results |
| `ZERO_RESULTS` | No results found | Handle as invalid ZIP |
| `OVER_QUERY_LIMIT` | Quota exceeded | Implement retry/backoff |
| `REQUEST_DENIED` | API key invalid | Check API key |
| `INVALID_REQUEST` | Invalid parameters | Fix request format |
| `UNKNOWN_ERROR` | Server error | Retry request |

## Usage Examples

### Basic ZIP Code Lookup
```bash
curl "https://maps.googleapis.com/maps/api/geocode/json?address=95401&components=country:US&key=YOUR_API_KEY"
```

### ZIP Code with Language Preference
```bash
curl "https://maps.googleapis.com/maps/api/geocode/json?address=95401&components=country:US&language=es&key=YOUR_API_KEY"
```

### Multiple ZIP Codes (Batch Processing)
```bash
for zip in 95401 90210 10001; do
  curl "https://maps.googleapis.com/maps/api/geocode/json?address=$zip&components=country:US&key=YOUR_API_KEY"
done
```

## Implementation in Petition System

### Area Scope Validation
The petition system uses address components to validate geographic scope:

```javascript
// State validation example
if (areaScope.scope === "state" && areaScope.name === "California") {
  const stateComponent = result.address_components.find(
    comp => comp.types.includes("administrative_area_level_1")
  );
  return stateComponent?.long_name === "California";
}
```

### City Extraction
```javascript
// Extract city name from geocoding result
const cityComponent = result.address_components.find(
  comp => comp.types.includes("locality")
);
const cityName = cityComponent?.long_name || "";
```

### Coordinate Extraction
```javascript
// Extract coordinates for mapping
const { lat, lng } = result.geometry.location;
```

## Rate Limits & Quotas

| Metric | Limit | Description |
|--------|-------|-------------|
| Requests per second | 50 | Standard Google Maps API |
| Daily quota | Variable | Based on API plan |
| Free tier | 2,500 requests/day | Google Maps Platform |
| Paid tier | Up to 100,000+ | Based on billing |

## Best Practices

### Request Optimization
- Use `components=country:US` to restrict to US locations
- Cache results for frequently requested ZIP codes
- Implement exponential backoff for rate limit errors
- Batch requests where possible

### Error Handling
```javascript
async function geocodeZip(zipCode) {
  try {
    const response = await fetch(
      `https://maps.googleapis.com/maps/api/geocode/json?address=${zipCode}&components=country:US&key=${API_KEY}`
    );
    const data = await response.json();
    
    if (data.status === 'OK') {
      return data.results[0];
    } else if (data.status === 'ZERO_RESULTS') {
      throw new Error('ZIP code not found');
    } else {
      throw new Error(`Geocoding error: ${data.status}`);
    }
  } catch (error) {
    console.error('Geocoding failed:', error);
    throw error;
  }
}
```

### Security Considerations
- Never expose API keys in client-side code
- Use server-side API key management
- Implement request rate limiting
- Monitor usage for abuse patterns
- Use environment variables for API keys

## Cost Management

### Monitoring Usage
- Track daily/monthly request volumes
- Monitor cost per API call
- Set up billing alerts
- Analyze usage patterns for optimization

### Optimization Strategies
- Implement aggressive caching
- Use component filtering to reduce irrelevant results
- Consider alternative data sources for high-volume ZIP lookups
- Batch geocoding requests when possible

## Integration Notes

### Response Processing
- Always check `status` field before processing results
- Handle `ZERO_RESULTS` gracefully
- Validate required address components exist
- Consider timezone implications for coordinate data

### Data Quality
- Google Maps data is generally accurate but may have delays
- Rural areas may have less precise geocoding
- New ZIP codes may not be immediately available
- Consider data freshness requirements

## Alternatives & Supplements

- **USPS Address Validation**: Official postal service data
- **Census Geocoder**: Free US government geocoding
- **OpenStreetMap/Nominatim**: Open source alternative
- **Commercial ZIP databases**: High-performance static data

## Related Implementation

See [`zip-to-city.go`](./zip-to-city.go) for the Go implementation that uses this API with area scope validation for the petition system.
