# Google Maps Geocoding API Documentation

## Overview
The Google Maps Geocoding API converts addresses (including ZIP codes) into geographic coordinates and structured address components.

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

## Further Reading

For complete documentation, see the [Official Google Geocoding API Documentation](https://developers.google.com/maps/documentation/geocoding/overview).
