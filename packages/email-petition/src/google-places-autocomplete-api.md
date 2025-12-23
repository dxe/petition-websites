# Google Places Autocomplete API Documentation

## Overview

The Google Places Autocomplete API provides real-time place predictions based on partial user input.

**Note**: The petition system uses the NEW Google Places API (v1) with POST requests and JSON body, not the legacy API with GET parameters.

## Response Structure

### Successful Response (New API Format)

```json
{
  "suggestions": [
    {
      "placePrediction": {
        "text": {
          "text": "Santa Rosa, CA, USA",
          "languageCode": "en"
        },
        "structuredFormat": {
          "mainText": {
            "text": "Santa Rosa"
          }
        }
      }
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

## Further Reading

For complete documentation, see the [Official Google Places API Documentation](https://developers.google.com/maps/documentation/places/web-service/autocomplete).
