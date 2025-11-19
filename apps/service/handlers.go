package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/dxe/service/config"
	"github.com/dxe/service/data"
	"github.com/dxe/service/model"
)

type CreateMessageInput struct {
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
	var body CreateMessageInput

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
