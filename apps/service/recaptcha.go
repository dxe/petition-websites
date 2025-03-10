package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/dxe/service/config"
)

type RecaptchaResponse struct {
	Success bool `json:"success"`
}

const verifyUrl = "https://www.google.com/recaptcha/api/siteverify"

func verifyRecaptcha(token string) (bool, error) {
	resp, err := http.PostForm(verifyUrl, url.Values{
		"secret":   {config.RecaptchaSecret},
		"response": {token},
	})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	var result RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}
	return result.Success, nil
}
