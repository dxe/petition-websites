"use client";

import { useState, useEffect, useRef } from "react";
import { Input } from "@dxe/petitions-components/input";

interface CityPrediction {
  cityAddress: string;
  cityName: string;
}

interface Coordinates {
  lat: number;
  lng: number;
}

interface CityAutocompleteProps {
  value: string;
  onChange: (value: string) => void;
  onBlur: () => void;
  disabled: boolean;
  placeholder?: string;
  // Coordinates at the epicenter of the primary city affiliated to the zipcode based on Google Geocoding API.
  // Used as the center of a custom search radius to restrict the cities that show up in the autocomplete dropdown.
  primaryCityCenterCoordinates: Coordinates;
  googleMapsPlacesApiKey: string;
}

export function CityAutocomplete({
  value,
  onChange,
  onBlur,
  disabled = false,
  placeholder = "Santa Rosa",
  primaryCityCenterCoordinates,
  googleMapsPlacesApiKey,
}: CityAutocompleteProps) {
  if (!googleMapsPlacesApiKey) {
    console.error("Google Maps Places API key not provided");
    return null;
  }

  const [predictions, setPredictions] = useState<CityPrediction[]>([]);
  const [isOpen, setIsOpen] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);
  const dropdownRef = useRef<HTMLDivElement>(null);

  const fetchPredictions = async (input: string) => {
    if (
      input.length == 0 ||
      !primaryCityCenterCoordinates.lat ||
      !primaryCityCenterCoordinates.lng
    ) {
      setPredictions([]);
      setIsOpen(false);
      return;
    }

    try {
      // Generate session token for cost optimization (max allowed length: 36 characters)
      const sessionToken = `ac-${Date.now().toString(36)}-${Math.random().toString(36).substring(2, 5)}`;

      // Request body for Google Places API
      const requestBody = {
        input,
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

      const response = await fetch(
        `https://places.googleapis.com/v1/places:autocomplete?key=${googleMapsPlacesApiKey}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "X-Goog-FieldMask":
              "suggestions.placePrediction.text,suggestions.placePrediction.structuredFormat.mainText.text",
          },
          body: JSON.stringify(requestBody),
        },
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();

      // Convert Google API response to our format
      const predictions: CityPrediction[] =
        data.suggestions?.map((suggestion: any) => {
          const cityAddress = suggestion.placePrediction?.text?.text || "";
          const cityName =
            suggestion.placePrediction?.structuredFormat?.mainText?.text ||
            cityAddress;

          return {
            cityAddress,
            cityName,
          };
        }) || [];

      setPredictions(predictions);
      setIsOpen(predictions.length > 0);
    } catch (error) {
      console.error("Error fetching city predictions:", error);
      setPredictions([]);
      setIsOpen(false);
    }
  };

  useEffect(() => {
    const timer = setTimeout(() => {
      if (value) {
        fetchPredictions(value);
      } else {
        setPredictions([]);
        setIsOpen(false);
      }
    }, 300);

    return () => clearTimeout(timer);
  }, [
    value,
    primaryCityCenterCoordinates.lat,
    primaryCityCenterCoordinates.lng,
  ]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node) &&
        inputRef.current &&
        !inputRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const selectPrediction = (prediction: CityPrediction) => {
    onChange(prediction.cityName);
    setPredictions([]);
    setIsOpen(false);
  };

  return (
    <div className="relative w-full">
      <Input
        ref={inputRef}
        type="text"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onBlur={onBlur}
        disabled={disabled}
        placeholder={placeholder}
        autoComplete="off"
      />
      {isOpen && predictions.length > 0 && (
        <div
          ref={dropdownRef}
          className="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-y-auto"
        >
          {predictions.map((prediction) => (
            <div
              key={prediction.cityName}
              className="px-3 py-2 cursor-pointer hover:bg-gray-100"
              onClick={() => selectPrediction(prediction)}
            >
              {prediction.cityAddress}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
