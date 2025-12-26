package geocoding

import (
	"fmt"
)

type AreaScope struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
}

type CityResult struct {
	Name          string  `json:"name"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	IsCityInScope bool    `json:"isCityInScope"`
}

// GetCityByZipCode uses Google Maps Geocoding API to convert ZIP codes to city names and coordinates
// See [google-geocoding-api.md](./google-geocoding-api.md) for complete API documentation
func GetCityByZipCode(zipCode string, areaScope *AreaScope, GoogleGeocodingAPIKey string) (*CityResult, error) {
	if GoogleGeocodingAPIKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	geocodingResp, err := GoogleGeocodingApiRequest(zipCode, GoogleGeocodingAPIKey)
	if err != nil {
		return nil, err
	}

	if geocodingResp.Status != "OK" {
		if geocodingResp.Status == "ZERO_RESULTS" {
			return nil, nil
		}
		return nil, fmt.Errorf("google geocoding API returned status: %s", geocodingResp.Status)
	}

	return extractCityFromGeocodingResponse(geocodingResp, zipCode, areaScope)
}
