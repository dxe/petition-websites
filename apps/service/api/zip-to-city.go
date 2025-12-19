package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// GoogleGeocodingResponse represents the response from Google Geocoding API
type GoogleGeocodingResponse struct {
	Results []GeocodingResult `json:"results"`
	Status  string            `json:"status"`
}

// GeocodingResult represents a single geocoding result
type GeocodingResult struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Coordinates       Coordinates        `json:"geometry"`
}

// Coordinates represents the coordinates of a geocoding result
type Coordinates struct {
	Location Location `json:"location"`
}

// Location represents latitude and longitude
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// AddressComponent represents an address component
type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// GetCityByZipCode uses Google Geocoding API to find the city name and coordinates from a ZIP code
func GetCityByZipCode(zipCode string) (string, float64, float64, error) {
	// Get API key from environment
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	fmt.Printf("[DEBUG] Loaded API key from environment: %s\n", apiKey)
	if apiKey == "" {
		fmt.Println("[DEBUG] GOOGLE_MAPS_API_KEY environment variable not set")
		return "", 0, 0, fmt.Errorf("GOOGLE_MAPS_API_KEY environment variable not set")
	}
	// Build the request URL
	baseURL := "https://maps.googleapis.com/maps/api/geocode/json"

	// URL encode the ZIP code
	params := url.Values{}
	params.Add("address", zipCode)
	params.Add("components", "country:US")
	params.Add("key", apiKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Printf("[DEBUG] Making request to URL: %s\n", fullURL)

	// Make HTTP request
	fmt.Printf("[DEBUG] Making HTTP GET request to Google Geocoding API\n")
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("[DEBUG] HTTP request failed: %v\n", err)
		return "", 0, 0, fmt.Errorf("failed to make request to Google Geocoding API: %w", err)
	}
	fmt.Printf("[DEBUG] HTTP response status: %s\n", resp.Status)
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, 0, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var geocodingResp GoogleGeocodingResponse
	err = json.Unmarshal(body, &geocodingResp)
	if err != nil {
		return "", 0, 0, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Check API response status
	fmt.Printf("[DEBUG] Google API response status: %s\n", geocodingResp.Status)
	if geocodingResp.Status != "OK" {
		if geocodingResp.Status == "ZERO_RESULTS" {
			fmt.Printf("[DEBUG] No results found for ZIP code: %s\n", zipCode)
			return "", 0, 0, fmt.Errorf("no results found for ZIP code: %s", zipCode)
		}
		fmt.Printf("[DEBUG] API returned non-OK status: %s\n", geocodingResp.Status)
		return "", 0, 0, fmt.Errorf("google geocoding API returned status: %s", geocodingResp.Status)
	}

	// Extract city from the first result
	fmt.Printf("[DEBUG] Found %d results in geocoding response\n", len(geocodingResp.Results))
	if len(geocodingResp.Results) == 0 {
		fmt.Printf("[DEBUG] No results array found for ZIP code: %s\n", zipCode)
		return "", 0, 0, fmt.Errorf("no results found for ZIP code: %s", zipCode)
	}

	// Find the city (locality) in address components
	fmt.Printf("[DEBUG] Searching for locality in %d address components\n", len(geocodingResp.Results[0].AddressComponents))
	for _, component := range geocodingResp.Results[0].AddressComponents {
		fmt.Printf("[DEBUG] Checking component: %s with types: %v\n", component.LongName, component.Types)
		for _, componentType := range component.Types {
			if componentType == "locality" {
				fmt.Printf("[DEBUG] Found locality: %s\n", component.LongName)
				lat := geocodingResp.Results[0].Coordinates.Location.Lat
				lng := geocodingResp.Results[0].Coordinates.Location.Lng
				return component.LongName, lat, lng, nil
			}
		}
	}

	// If no locality found, try administrative_area_level_2 (county)
	for _, component := range geocodingResp.Results[0].AddressComponents {
		for _, componentType := range component.Types {
			if componentType == "administrative_area_level_2" {
				lat := geocodingResp.Results[0].Coordinates.Location.Lat
				lng := geocodingResp.Results[0].Coordinates.Location.Lng
				return component.LongName, lat, lng, nil
			}
		}
	}

	// If still no city found, return formatted address as fallback
	if len(geocodingResp.Results) > 0 {
		lat := geocodingResp.Results[0].Coordinates.Location.Lat
		lng := geocodingResp.Results[0].Coordinates.Location.Lng
		return geocodingResp.Results[0].FormattedAddress, lat, lng, nil
	}

	return "", 0, 0, fmt.Errorf("could not extract city name from geocoding response for ZIP code: %s", zipCode)
}
