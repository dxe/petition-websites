"use client";

import { useState, useEffect, useRef } from "react";
import { Input } from "@dxe/petitions-components/input";
import ky from "ky";

interface CityPrediction {
  cityAddress: string;
  cityName: string;
}

interface CityAutocompleteProps {
  value: string;
  onChange: (value: string) => void;
  onBlur: () => void;
  disabled?: boolean;
  placeholder?: string;
  lat?: number;
  lng?: number;
}

export function CityAutocomplete({
  value,
  onChange,
  onBlur,
  disabled = false,
  placeholder = "Santa Rosa",
  lat,
  lng,
}: CityAutocompleteProps) {
  const [predictions, setPredictions] = useState<CityPrediction[]>([]);
  const [isOpen, setIsOpen] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);
  const dropdownRef = useRef<HTMLDivElement>(null);

  const fetchPredictions = async (input: string) => {
    if (input.length == 0 || !lat || !lng) {
      setPredictions([]);
      setIsOpen(false);
      return;
    }

    try {
      const response = await ky
        .post(
          `${process.env.NEXT_PUBLIC_CAMPAIGN_MAILER_API_ROOT}/cityAutocomplete`,
          {
            json: { input, lat, lng },
          },
        )
        .json<{ predictions: CityPrediction[] }>();

      setPredictions(response.predictions);
      setIsOpen(response.predictions.length > 0);
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
  }, [value, lat, lng]);

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
