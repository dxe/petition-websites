package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// PlacesAutocompleteRequest represents request for Places Autocomplete API
type PlacesAutocompleteRequest struct {
	Input  string  `json:"input"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Radius int     `json:"radius"` // in meters
	Types  string  `json:"types"`  // restrict to cities
}

// PlacesAutocompleteResponse represents response from Places Autocomplete API
type PlacesAutocompleteResponse struct {
	Predictions []PlacePrediction `json:"predictions"`
	Status      string            `json:"status"`
}

// PlacePrediction represents a single place prediction
type PlacePrediction struct {
	Description string   `json:"description"`
	PlaceId     string   `json:"place_id"`
	Types       []string `json:"types"`
}

// CityPrediction represents simplified city prediction response
type CityPrediction struct {
	City    string `json:"city"`
	PlaceId string `json:"place_id"`
}

// GetCityPredictions uses Google Places Autocomplete API to find city predictions near coordinates
func GetCityPredictions(input string, lat, lng float64) ([]CityPrediction, error) {
	// Get API key from environment
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	fmt.Printf("[DEBUG] Loaded API key for Places API: %s\n", apiKey)
	if apiKey == "" {
		fmt.Println("[DEBUG] GOOGLE_MAPS_API_KEY environment variable not set")
		return nil, fmt.Errorf("GOOGLE_MAPS_API_KEY environment variable not set")
	}

	// Build the request URL
	baseURL := "https://maps.googleapis.com/maps/api/place/autocomplete/json"

	// URL encode the parameters
	params := url.Values{}
	params.Add("input", input)
	params.Add("location", fmt.Sprintf("%f,%f", lat, lng))
	params.Add("radius", "5000") // 5km in meters
	params.Add("types", "(cities)")
	params.Add("key", apiKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Printf("[DEBUG] Making Places API request to URL: %s\n", fullURL)

	// Make HTTP request
	fmt.Printf("[DEBUG] Making HTTP GET request to Google Places API\n")
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("[DEBUG] Places API HTTP request failed: %v\n", err)
		return nil, fmt.Errorf("failed to make request to Google Places API: %w", err)
	}
	fmt.Printf("[DEBUG] Places API HTTP response status: %s\n", resp.Status)
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var placesResp PlacesAutocompleteResponse
	err = json.Unmarshal(body, &placesResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Places API JSON response: %w", err)
	}

	// Check API response status
	fmt.Printf("[DEBUG] Places API response status: %s\n", placesResp.Status)
	if placesResp.Status != "OK" {
		if placesResp.Status == "ZERO_RESULTS" {
			fmt.Printf("[DEBUG] No predictions found for input: %s\n", input)
			return []CityPrediction{}, nil
		}
		fmt.Printf("[DEBUG] Places API returned non-OK status: %s\n", placesResp.Status)
		return nil, fmt.Errorf("google Places API returned status: %s", placesResp.Status)
	}

	// Convert to simplified predictions
	fmt.Printf("[DEBUG] Found %d predictions\n", len(placesResp.Predictions))
	var predictions []CityPrediction
	for _, prediction := range placesResp.Predictions {
		// Extract city name from description (usually format: "City, State, Country")
		cityName := prediction.Description
		if commaIndex := strings.Index(cityName, ","); commaIndex != -1 {
			cityName = cityName[:commaIndex]
		}

		predictions = append(predictions, CityPrediction{
			City:    cityName,
			PlaceId: prediction.PlaceId,
		})
		fmt.Printf("[DEBUG] Prediction: %s (PlaceId: %s)\n", cityName, prediction.PlaceId)
	}

	return predictions, nil
}
