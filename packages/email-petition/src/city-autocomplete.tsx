"use client";

import { useState, useEffect, useRef } from "react";
import { Input } from "@dxe/petitions-components/input";
import {
  googleMapsPlacesAutoCompleteApiResponse,
  CityPrediction,
  Coordinates,
} from "./google-maps-places-autocomplete-api";

interface CityAutocompleteProps {
  searchInput: string;
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
  searchInput,
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

  const fetchCityPredictions = async (
    searchQuery: string,
  ): Promise<CityPrediction[]> => {
    if (
      searchQuery.length == 0 ||
      !primaryCityCenterCoordinates.lat ||
      !primaryCityCenterCoordinates.lng
    ) {
      return [];
    }

    try {
      return await googleMapsPlacesAutoCompleteApiResponse(
        searchQuery,
        primaryCityCenterCoordinates,
        googleMapsPlacesApiKey,
      );
    } catch (error) {
      console.error("Error fetching city predictions:", error);
      return [];
    }
  };

  const updatePredictionsState = (predictions: CityPrediction[]) => {
    setPredictions(predictions);
    setIsOpen(predictions.length > 0);
  };

  const clearPredictionsState = () => {
    setPredictions([]);
    setIsOpen(false);
  };

  useEffect(() => {
    const timer = setTimeout(async () => {
      if (searchInput) {
        const predictions = await fetchCityPredictions(searchInput);
        updatePredictionsState(predictions);
      } else {
        clearPredictionsState();
      }
    }, 300);

    return () => clearTimeout(timer);
  }, [
    searchInput,
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
    clearPredictionsState();
  };

  return (
    <div className="relative w-full">
      <Input
        ref={inputRef}
        type="text"
        value={searchInput}
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
