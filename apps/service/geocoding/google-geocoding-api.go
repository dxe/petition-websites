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

// GoogleGeocodingApiRequest makes a request to the Google Geocoding API
func GoogleGeocodingApiRequest(zipCode string, GoogleGeocodingAPIKey string) (*GoogleGeocodingResponse, error) {
	baseURL := "https://maps.googleapis.com/maps/api/geocode/json"

	params := url.Values{}
	params.Add("address", zipCode)
	params.Add("components", "country:US")
	params.Add("key", GoogleGeocodingAPIKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to Google Geocoding API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var geocodingResp GoogleGeocodingResponse
	err = json.Unmarshal(body, &geocodingResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &geocodingResp, nil
}
