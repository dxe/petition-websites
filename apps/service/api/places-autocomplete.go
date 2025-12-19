package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// PlacesAutocompleteRequest represents request for Places Autocomplete API
type PlacesAutocompleteRequest struct {
	Input  string  `json:"input"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Radius int     `json:"radius"` // in meters
	Types  string  `json:"types"`  // restrict to cities
}

// PlacesAutocompleteResponse represents response from NEW Places Autocomplete API
type PlacesAutocompleteResponse struct {
	Suggestions []PlaceSuggestion `json:"suggestions"`
}

// PlaceSuggestion represents a single place suggestion from new API
type PlaceSuggestion struct {
	PlacePrediction PlacePrediction `json:"placePrediction"`
}

// PlacePrediction represents a single place prediction
type PlacePrediction struct {
	Text             Text             `json:"text"`
	StructuredFormat StructuredFormat `json:"structuredFormat"`
}

// StructuredFormat represents structured format with main and secondary text
type StructuredFormat struct {
	MainText Text `json:"mainText"`
}

// Text represents formatted text with highlighted substrings
type Text struct {
	Text string `json:"text"`
}

// CityPrediction represents simplified city prediction response
type CityPrediction struct {
	CityAddress string `json:"cityAddress"`
	CityName    string `json:"cityName"`
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

	// Build the request URL for NEW Places API
	baseURL := "https://places.googleapis.com/v1/places:autocomplete"

	// Request body for NEW Places API
	requestBody := map[string]interface{}{
		"input": input,
		"locationRestriction": map[string]interface{}{
			"circle": map[string]interface{}{
				"center": map[string]float64{"latitude": lat, "longitude": lng},
				"radius": 50000,
			},
		},
		"includedPrimaryTypes": []string{"locality"},
	}

	// Make HTTP request
	fmt.Printf("[DEBUG] Making HTTP POST request to Google Places API\n")

	// Marshal request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create HTTP request with API key in header
	req, err := http.NewRequest("POST", baseURL+"?key="+apiKey, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-FieldMask", "suggestions.placePrediction.text,suggestions.placePrediction.structuredFormat.mainText.text")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
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

	// Check if we have suggestions
	fmt.Printf("[DEBUG] Found %d suggestions\n", len(placesResp.Suggestions))
	if len(placesResp.Suggestions) == 0 {
		fmt.Printf("[DEBUG] No predictions found for input: %s\n", input)
		return []CityPrediction{}, nil
	}

	// Convert to simplified predictions
	var predictions []CityPrediction
	for _, suggestion := range placesResp.Suggestions {
		cityAddress := suggestion.PlacePrediction.Text.Text
		cityName := suggestion.PlacePrediction.StructuredFormat.MainText.Text

		predictions = append(predictions, CityPrediction{
			CityAddress: cityAddress,
			CityName:    cityName,
		})
		fmt.Printf("[DEBUG] Prediction - CityName: %s, CityAddress: %s\n", cityName, cityAddress)
	}

	return predictions, nil
}
