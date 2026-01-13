package geocoding

import "fmt"

func extractCityFromGeocodingResponse(geocodingResp *GoogleGeocodingResponse, zipCode string, areaScope *AreaScope) (*CityResult, error) {

	if len(geocodingResp.Results) == 0 {
		return nil, fmt.Errorf("no results found for ZIP code: %s", zipCode)
	}

	isCityInScope := true

	// Validate area scope if provided
	if areaScope != nil {

		// Custom defined area scope to Google Maps component type
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

		// Find the matching scope component in API response and validate
		scopeMatched := false
		for _, component := range geocodingResp.Results[0].AddressComponents {
			for _, ct := range component.Types {
				if ct == componentType {
					if component.LongName == areaScope.Name {
						scopeMatched = true
						break
					} else {
						isCityInScope = false
						scopeMatched = true
						break
					}
				}
			}
			if scopeMatched {
				break
			}
		}

		if !scopeMatched {
			isCityInScope = false
		}
	}

	// Extract address components
	var city, state string
	for _, component := range geocodingResp.Results[0].AddressComponents {
		for _, componentType := range component.Types {
			switch componentType {
			case "locality":
				city = component.LongName
			case "administrative_area_level_1":
				// Two letter state code
				state = component.ShortName
			}
		}
	}

	if city != "" {
		lat := geocodingResp.Results[0].Coordinates.Location.Lat
		lng := geocodingResp.Results[0].Coordinates.Location.Lng
		return &CityResult{
			City:          city,
			State:         state,
			Latitude:      lat,
			Longitude:     lng,
			IsCityInScope: isCityInScope,
		}, nil
	}

	return nil, nil
}
