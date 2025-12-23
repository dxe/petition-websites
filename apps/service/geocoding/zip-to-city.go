package geocoding

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GoogleGeocodingResponse struct {
	Results []GeocodingResult `json:"results"`
	Status  string            `json:"status"`
}

type GeocodingResult struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Coordinates       Coordinates        `json:"geometry"`
}

type Coordinates struct {
	Location Location `json:"location"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type AreaScope struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
}

type CityResult struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GetCityByZipCode uses Google Maps Geocoding API to convert ZIP codes to city names and coordinates
// See [google-geocoding-api.md](./google-geocoding-api.md) for complete API documentation
func GetCityByZipCode(zipCode string, areaScope *AreaScope, GoogleGeocodingAPIKey string) (*CityResult, error) {
	if GoogleGeocodingAPIKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}
	// Build the request URL
	baseURL := "https://maps.googleapis.com/maps/api/geocode/json"

	// URL encode the ZIP code
	params := url.Values{}
	params.Add("address", zipCode)
	params.Add("components", "country:US")
	params.Add("key", GoogleGeocodingAPIKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make HTTP request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to Google Geocoding API: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var geocodingResp GoogleGeocodingResponse
	err = json.Unmarshal(body, &geocodingResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Check API response status
	if geocodingResp.Status != "OK" {
		if geocodingResp.Status == "ZERO_RESULTS" {
			return nil, fmt.Errorf("no results found for ZIP code: %s", zipCode)
		}
		return nil, fmt.Errorf("google geocoding API returned status: %s", geocodingResp.Status)
	}

	// Extract city from the first result
	if len(geocodingResp.Results) == 0 {
		return nil, fmt.Errorf("no results found for ZIP code: %s", zipCode)
	}

	// Validate area scope if provided
	if areaScope != nil {
		// Map scope to Google Maps component type
		var componentType string
		switch areaScope.Scope {
		case "city":
			componentType = "locality"
		case "county":
			componentType = "administrative_area_level_2"
		case "state":
			componentType = "administrative_area_level_1"
		case "country":
			componentType = "country"
		default:
			return nil, fmt.Errorf("unknown area scope: %s", areaScope.Scope)
		}

		// Find the matching component and validate
		scopeMatched := false
		for _, component := range geocodingResp.Results[0].AddressComponents {
			for _, ct := range component.Types {
				if ct == componentType {
					if component.LongName == areaScope.Name {
						scopeMatched = true
						break
					} else {
						return nil, fmt.Errorf("area scope validation failed: expected %s, got %s", areaScope.Name, component.LongName)
					}
				}
			}
			if scopeMatched {
				break
			}
		}

		if !scopeMatched {
			return nil, fmt.Errorf("no matching component found for area scope: %s", componentType)
		}
	}

	// Find the city (locality) in address components
	for _, component := range geocodingResp.Results[0].AddressComponents {
		for _, componentType := range component.Types {
			if componentType == "locality" {
				lat := geocodingResp.Results[0].Coordinates.Location.Lat
				lng := geocodingResp.Results[0].Coordinates.Location.Lng
				return &CityResult{
					Name:      component.LongName,
					Latitude:  lat,
					Longitude: lng,
				}, nil
			}
		}
	}

	// If no locality found, try administrative_area_level_2 (county)
	for _, component := range geocodingResp.Results[0].AddressComponents {
		for _, componentType := range component.Types {
			if componentType == "administrative_area_level_2" {
				lat := geocodingResp.Results[0].Coordinates.Location.Lat
				lng := geocodingResp.Results[0].Coordinates.Location.Lng
				return &CityResult{
					Name:      component.LongName,
					Latitude:  lat,
					Longitude: lng,
				}, nil
			}
		}
	}

	// If still no city found, return formatted address as fallback
	if len(geocodingResp.Results) > 0 {
		lat := geocodingResp.Results[0].Coordinates.Location.Lat
		lng := geocodingResp.Results[0].Coordinates.Location.Lng
		return &CityResult{
			Name:      geocodingResp.Results[0].FormattedAddress,
			Latitude:  lat,
			Longitude: lng,
		}, nil
	}

	return nil, fmt.Errorf("could not extract city name from geocoding response for ZIP code: %s", zipCode)
}
