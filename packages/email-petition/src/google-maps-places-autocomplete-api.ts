// Google Maps Places API types and HTTP transport functions

export interface GoogleMapsPlacesRequest {
  input: string;
  locationRestriction: {
    circle: {
      center: {
        latitude: number;
        longitude: number;
      };
      radius: number;
    };
  };
  includedPrimaryTypes: ["locality"];
  sessionToken: string;
}

export interface GoogleMapsPlacesResponse {
  suggestions?: GoogleMapsSuggestion[];
}

export interface GoogleMapsSuggestion {
  placePrediction: {
    text: {
      text: string;
    };
    structuredFormat: {
      mainText: {
        text: string;
      };
    };
  };
}

export interface CityPrediction {
  cityAddress: string;
  cityName: string;
}

export interface Coordinates {
  lat: number;
  lng: number;
}

export async function googleMapsPlacesAutoCompleteApiResponse(
  searchQuery: string,
  primaryCityCenterCoordinates: Coordinates,
  googleMapsPlacesApiKey: string,
): Promise<CityPrediction[]> {
  // Generate session token for cost optimization (max allowed length: 36 characters)
  const sessionToken = `ac-${Date.now().toString(36)}-${Math.random().toString(36).substring(2, 5)}`;

  const requestBody: GoogleMapsPlacesRequest = {
    input: searchQuery,
    locationRestriction: {
      circle: {
        center: {
          latitude: primaryCityCenterCoordinates.lat,
          longitude: primaryCityCenterCoordinates.lng,
        },
        // 50 km is the maximum allowed by the autocomplete API: https://developers.google.com/maps/documentation/javascript/reference/places-service#LocationRestriction
        // Note: This may exclude some relevant cities in dispersed regions.
        radius: 50000,
      },
    },
    includedPrimaryTypes: ["locality"],
    sessionToken,
  };

  const url = `https://places.googleapis.com/v1/places:autocomplete?key=${googleMapsPlacesApiKey}`;

  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Goog-FieldMask":
        "suggestions.placePrediction.text,suggestions.placePrediction.structuredFormat.mainText.text",
    },
    body: JSON.stringify(requestBody),
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const data: GoogleMapsPlacesResponse = await response.json();

  const predictions: CityPrediction[] =
    data.suggestions?.map((suggestion: GoogleMapsSuggestion) => {
      const cityAddress = suggestion.placePrediction?.text?.text || "";
      const cityName =
        suggestion.placePrediction?.structuredFormat?.mainText?.text ||
        cityAddress;

      return {
        cityAddress,
        cityName,
      };
    }) || [];

  return predictions;
}
