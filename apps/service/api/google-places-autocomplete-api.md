# Google Places Autocomplete API Documentation

## Overview
The Google Places Autocomplete API provides real-time place predictions based on partial user input. This service is used in the petition system to provide city name suggestions as users type, improving form usability and data accuracy.

**Note**: The petition system uses the NEW Google Places API (v1) with POST requests and JSON body, not the legacy API with GET parameters.

## API Endpoint
```
POST https://places.googleapis.com/v1/places:autocomplete
```

## Request Headers

```
Content-Type: application/json
X-Goog-Api-Key: YOUR_API_KEY
X-Goog-FieldMask: places.displayName,places.formattedAddress,places.types
```

## Request Body

### Basic Request Structure
```json
{
  "input": "Santa",
  "includedPrimaryTypes": ["locality"]
}
```

### Request with Location Restriction (Used in getCityPredictions)
```json
{
  "input": "Santa",
  "locationRestriction": {
    "circle": {
      "center": {
        "latitude": 38.4405,
        "longitude": -122.7144
      },
      "radius": 50000
    }
  },
  "includedPrimaryTypes": ["locality"]
}
```

### Request Body Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `input` | string | Yes | Partial search term (minimum 1 character) |
| `locationRestriction` | object | No | Geographic area to restrict results |
| `locationRestriction.circle` | object | Yes (if locationRestriction used) | Circular search area |
| `locationRestriction.circle.center` | object | Yes | Center point coordinates |
| `locationRestriction.circle.center.latitude` | number | Yes | Latitude of center point |
| `locationRestriction.circle.center.longitude` | number | Yes | Longitude of center point |
| `locationRestriction.circle.radius` | number | Yes | Search radius in meters (max 50000) |
| `includedPrimaryTypes` | array | No | Filter by place types |
| `languageCode` | string | No | Response language (default: English) |
| `regionCode` | string | No | Region code (e.g., "us") |

### Parameters Used in getCityPredictions

The `getCityPredictions` function specifically uses these parameters:

```go
requestBody := map[string]interface{}{
    "input": input,                                    // User's partial input
    "locationRestriction": map[string]interface{}{     // Geographic restriction
        "circle": map[string]interface{}{
            "center": map[string]float64{              // Center coordinates
                "latitude": lat,                       // From user's ZIP/location
                "longitude": lng,
            },
            "radius": 50000,                           // 50km search radius
        },
    },
    "includedPrimaryTypes": []string{"locality"},     // Cities only
}
```

## Response Structure

### Successful Response (New API Format)
```json
{
  "places": [
    {
      "displayName": {
        "text": "Santa Rosa",
        "languageCode": "en"
      },
      "formattedAddress": "Santa Rosa, CA, USA",
      "types": ["locality", "political"],
      "placeId": "ChIJr8p0nB2mj4ARP5yVCLLkQzA"
    },
    {
      "displayName": {
        "text": "Santa Monica",
        "languageCode": "en"
      },
      "formattedAddress": "Santa Monica, CA, USA", 
      "types": ["locality", "political"],
      "placeId": "ChIJ2_UUQFq2woARCAgVd2kG2xU"
    }
  ]
}
```

### Error Response
```json
{
  "error": {
    "code": 400,
    "message": "Invalid request",
    "status": "INVALID_ARGUMENT"
  }
}
```

## Place Structure (New API)

Each place object contains:

| Field | Type | Description |
|-------|------|-------------|
| `displayName.text` | string | Primary place name |
| `displayName.languageCode` | string | Language of the name |
| `formattedAddress` | string | Full formatted address |
| `types` | array | Place type classifications |
| `placeId` | string | Unique place identifier |

## Status Codes

| Status | Description | Action |
|--------|-------------|--------|
| `OK` | Request successful | Display predictions |
| `ZERO_RESULTS` | No predictions found | Show "no results" message |
| `OVER_QUERY_LIMIT` | Quota exceeded | Implement retry/backoff |
| `REQUEST_DENIED` | API key invalid | Check API key |
| `INVALID_REQUEST` | Invalid parameters | Fix request format |

## Usage Examples

### Basic City Autocomplete (New API)
```bash
curl -X POST "https://places.googleapis.com/v1/places:autocomplete" \
  -H "Content-Type: application/json" \
  -H "X-Goog-Api-Key: YOUR_API_KEY" \
  -H "X-Goog-FieldMask: places.displayName,places.formattedAddress,places.types" \
  -d '{
    "input": "Santa",
    "includedPrimaryTypes": ["locality"]
  }'
```

### Location-Restricted Autocomplete (As Used in getCityPredictions)
```bash
curl -X POST "https://places.googleapis.com/v1/places:autocomplete" \
  -H "Content-Type: application/json" \
  -H "X-Goog-Api-Key: YOUR_API_KEY" \
  -H "X-Goog-FieldMask: places.displayName,places.formattedAddress,places.types" \
  -d '{
    "input": "Santa",
    "locationRestriction": {
      "circle": {
        "center": {
          "latitude": 38.4405,
          "longitude": -122.7144
        },
        "radius": 50000
      }
    },
    "includedPrimaryTypes": ["locality"]
  }'
```

### Go Implementation Example (getCityPredictions)
```go
func GetCityPredictions(input string, lat, lng float64) ([]CityPrediction, error) {
    apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
    
    requestBody := map[string]interface{}{
        "input": input,
        "locationRestriction": map[string]interface{}{
            "circle": map[string]interface{}{
                "center": map[string]float64{
                    "latitude": lat,
                    "longitude": lng,
                },
                "radius": 50000,
            },
        },
        "includedPrimaryTypes": []string{"locality"},
    }
    
    // Convert to JSON and make POST request...
}
```

## Implementation in Petition System

### Request Structure Used in Service
The service uses the NEW Places API with these specific parameters:

```go
requestBody := map[string]interface{}{
    "input": input,                                    // User's partial input
    "locationRestriction": map[string]interface{}{     // Geographic restriction
        "circle": map[string]interface{}{
            "center": map[string]float64{              // Center coordinates
                "latitude": lat,                       // From ZIP geocoding
                "longitude": lng,
            },
            "radius": 50000,                           // 50km search radius
        },
    },
    "includedPrimaryTypes": []string{"locality"},     // Cities only
}
```

### Response Processing (Go)
```go
type PlaceResponse struct {
    Places []Place `json:"places"`
}

type Place struct {
    DisplayName struct {
        Text string `json:"text"`
    } `json:"displayName"`
    FormattedAddress string   `json:"formattedAddress"`
    Types           []string `json:"types"`
    PlaceId         string   `json:"placeId"`
}

// Convert to internal CityPrediction format
for _, place := range response.Places {
    predictions = append(predictions, CityPrediction{
        Description: place.FormattedAddress,
        PlaceId:     place.PlaceId,
        MainText:    place.DisplayName.Text,
        Types:       place.Types,
    })
}
```

## Related Implementation

See [`places-autocomplete.go`](./places-autocomplete.go) for the Go implementation that uses this API with location restriction and city filtering for the petition system's `getCityPredictions` function.

### Client-Side Integration
```javascript
// Debounced autocomplete for form input
let debounceTimer;
const cityInput = document.getElementById('city-input');

cityInput.addEventListener('input', (e) => {
  clearTimeout(debounceTimer);
  const input = e.target.value;
  
  if (input.length < 2) {
    clearSuggestions();
    return;
  }
  
  debounceTimer = setTimeout(() => {
    getCityPredictions(input, userLat, userLng)
      .then(suggestions => showSuggestions(suggestions));
  }, 300);
});
```

## Rate Limits & Quotas

| Metric | Limit | Description |
|--------|-------|-------------|
| Requests per second | 100 | Standard Google Places API |
| Daily quota | Variable | Based on API plan |
| Free tier | 1,000 requests/day | Google Maps Platform |
| Paid tier | Up to 100,000+ | Based on billing |

## Best Practices

### Request Optimization
- Use `types=(cities)` to filter for cities only
- Implement client-side debouncing (300ms recommended)
- Use location biasing for better local results
- Cache frequent predictions for performance

### Debouncing Implementation
```javascript
function debounce(func, delay) {
  let timeoutId;
  return function(...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => func.apply(this, args), delay);
  };
}

const debouncedAutocomplete = debounce(getCityPredictions, 300);
```

### Error Handling
```javascript
async function safeAutocomplete(input) {
  try {
    if (input.length < 2) return [];
    
    const response = await fetchAutocomplete(input);
    
    switch (response.status) {
      case 'OK':
        return response.predictions;
      case 'ZERO_RESULTS':
        return [];
      case 'OVER_QUERY_LIMIT':
        console.warn('Rate limit exceeded, retrying...');
        await new Promise(resolve => setTimeout(resolve, 1000));
        return safeAutocomplete(input);
      default:
        console.error(`Autocomplete error: ${response.status}`);
        return [];
    }
  } catch (error) {
    console.error('Autocomplete failed:', error);
    return [];
  }
}
```

### Security Considerations
- Never expose API keys in client-side code
- Use server-side proxy for API requests
- Implement request rate limiting per user/session
- Monitor for automated scraping attempts
- Use environment variables for API keys

## Cost Management

### Monitoring Usage
- Track autocomplete requests per user
- Monitor character input patterns
- Analyze most common search terms
- Set up billing alerts for quota limits

### Optimization Strategies
- Increase debounce delay for high-traffic scenarios
- Implement local caching of recent predictions
- Use location biasing to reduce irrelevant results
- Consider static city data for common locations

## Integration Notes

### UI/UX Considerations
- Show loading state during API requests
- Highlight matched portions in suggestions
- Implement keyboard navigation (arrow keys, enter)
- Handle empty states gracefully
- Consider accessibility for screen readers

### Data Quality
- Google Places data is generally comprehensive
- Rural areas may have fewer suggestions
- New cities may take time to appear
- Consider timezone and regional variations

### Performance Optimization
```javascript
// Client-side caching
const predictionCache = new Map();

async function getCachedPredictions(input) {
  if (predictionCache.has(input)) {
    return predictionCache.get(input);
  }
  
  const predictions = await getCityPredictions(input);
  predictionCache.set(input, predictions);
  
  // Clear cache after 5 minutes
  setTimeout(() => predictionCache.delete(input), 300000);
  
  return predictions;
}
```

## Alternatives & Supplements

- **Static City Database**: Pre-loaded city data for common locations
- **USPS City API**: Official postal service city data
- **OpenStreetMap Nominatim**: Open source autocomplete
- **Algolia Places**: Commercial autocomplete service
- **Mapbox Geocoding**: Alternative geocoding service

## Advanced Features

### Place Details Enhancement
```javascript
// Get detailed place information after selection
async function getPlaceDetails(placeId) {
  const url = `https://maps.googleapis.com/maps/api/place/details/json?place_id=${placeId}&fields=name,formatted_address,geometry&key=${API_KEY}`;
  
  const response = await fetch(url);
  const data = await response.json();
  
  if (data.status === 'OK') {
    return data.result;
  }
}
```

### Session Token Usage (Cost Optimization)
The refactored implementation now uses session tokens for cost optimization:

```go
// Generate session token for cost optimization
sessionToken := NewSessionToken()

// Include in request body
requestBody := map[string]interface{}{
    "input": input,
    "sessionToken": sessionToken,  // Reduces API costs
    // ... other parameters
}
```

**Benefits of Session Tokens:**
- **Cost Reduction**: Autocomplete + Place Details counts as one session
- **Better Performance**: Reduced API latency for follow-up requests
- **Quota Efficiency**: Optimized usage of Google Maps quota

**Implementation Details:**
- Session tokens are generated using `NewSessionToken()` function
- Tokens are included in autocomplete requests
- Same token used for subsequent Place Details API calls
- Tokens expire after 3 minutes of inactivity

### Go Implementation Example (Updated)
```go
func GetCityPredictions(input string, lat, lng float64) ([]CityPrediction, error) {
    // Generate session token for cost optimization
    sessionToken := NewSessionToken()
    
    requestBody := map[string]interface{}{
        "input": input,
        "sessionToken": sessionToken,  // Cost optimization
        "locationRestriction": map[string]interface{}{
            "circle": map[string]interface{}{
                "center": map[string]float64{
                    "latitude": lat,
                    "longitude": lng,
                },
                "radius": 50000,
            },
        },
        "includedPrimaryTypes": []string{"locality"},
    }
    
    // Make request with session token...
}
```
