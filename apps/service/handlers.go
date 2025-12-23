package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/dxe/service/config"
	"github.com/dxe/service/data"
	"github.com/dxe/service/geocoding"
	"github.com/dxe/service/model"
)

type CreateMessageReq struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
	OutsideUS bool   `json:"outside_us"`
	Zip       string `json:"zip,omitempty"`
	City      string `json:"city,omitempty"`
	Message   string `json:"message"`
	Token     string `json:"token"`
	Campaign  string `json:"campaign,omitempty"`
}

func (s *server) createMessageHandler(w http.ResponseWriter, r *http.Request) {
	var body CreateMessageReq

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	if config.RecaptchaSecret == "" {
		fmt.Println("Recaptcha secret not set, skipping verification")
	} else {
		ok, err := verifyRecaptcha(body.Token)
		if err != nil {
			fmt.Printf("error verifying recaptcha: %v\n", err)
			http.Error(w, "error verifying recaptcha", http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Error(w, "invalid captcha", http.StatusForbidden)
			return
		}
	}

	_, err = mail.ParseAddress(body.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid email address: %v", err), http.StatusBadRequest)
		return
	}

	if len(body.Name) == 0 {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	message := model.Message{
		Name:  body.Name,
		Email: body.Email,
		Phone: sql.NullString{
			String: body.Phone,
			Valid:  body.Phone != "",
		},
		OutsideUS: body.OutsideUS,
		Zip: sql.NullString{
			String: body.Zip,
			Valid:  body.Zip != "",
		},
		City: sql.NullString{
			String: body.City,
			Valid:  body.City != "",
		},
		Message: body.Message,
		IPAddress: sql.NullString{
			String: r.RemoteAddr,
			Valid:  r.RemoteAddr != "",
		},
		Campaign: sql.NullString{
			String: body.Campaign,
			Valid:  body.Campaign != "",
		},
	}

	err = model.InsertMessage(s.db, message)
	if err != nil {
		http.Error(w, fmt.Sprintf("error saving message: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ok"))
}

type GetTallyResp struct {
	Total int `json:"total"`
}

func (s *server) getTallyHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	campaign := queryParams.Get("campaign")
	if campaign == "" {
		http.Error(w, "missing campaign", http.StatusBadRequest)
		return
	}

	tally, err := model.GetTally(s.db, campaign)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting tally: %v", err), http.StatusInternalServerError)
		return
	}

	resp := GetTallyResp{
		Total: tally,
	}
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

type GetAssemblyMembersReq struct {
	City string `json:"city"`
	Zip  string `json:"zip"`
}

type GetAssemblyMembersResp struct {
	Members []data.AssemblyMember `json:"members"`
}

func (s *server) getAssemblyMembersHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	city := queryParams.Get("city")
	zipcode := queryParams.Get("zip")

	if city == "" || zipcode == "" {
		http.Error(w, "city and zipcode parameters are required", http.StatusBadRequest)
		return
	}

	membersMap := config.GetAssemblyMembers(data.Municipality(city), data.Zip(zipcode))
	members := make([]data.AssemblyMember, 0, len(membersMap))
	for _, member := range membersMap {
		members = append(members, member)
	}

	resp := GetAssemblyMembersResp{
		Members: members,
	}
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// ZipToCityLookupReq represents request for city lookup by ZIP code
type ZipToCityLookupReq struct {
	ZipCode   string `json:"zip_code"`
	AreaScope *struct {
		Name  string `json:"name"`
		Scope string `json:"scope"`
	} `json:"areaScope,omitempty"`
}

// ZipToCityLookupResp represents response for city lookup
type ZipToCityLookupResp struct {
	City string  `json:"city"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

// CityAutocompleteReq represents request for city autocomplete
type CityAutocompleteReq struct {
	Input string  `json:"input"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

// zipToCityLookupHandler handles ZIP code to city lookup requests
func (s *server) zipToCityLookupHandler(w http.ResponseWriter, r *http.Request) {
	var req ZipToCityLookupReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.ZipCode == "" {
		http.Error(w, "zip_code is required", http.StatusBadRequest)
		return
	}

	if config.GoogleMapsGeocodingAPIKey == "" {
		http.Error(w, "Google Maps Geocoding API key not configured", http.StatusInternalServerError)
		return
	}

	// Convert areaScope to API type if provided
	var areaScope *geocoding.AreaScope
	if req.AreaScope != nil {
		areaScope = &geocoding.AreaScope{
			Name:  req.AreaScope.Name,
			Scope: req.AreaScope.Scope,
		}
	}

	cityResult, err := geocoding.GetCityByZipCode(req.ZipCode, areaScope, config.GoogleMapsGeocodingAPIKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to lookup city: %v", err), http.StatusInternalServerError)
		return
	}

	city := cityResult.Name
	lat := cityResult.Latitude
	lng := cityResult.Longitude

	resp := ZipToCityLookupResp{
		City: city,
		Lat:  lat,
		Lng:  lng,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
