package config

import (
	"os"

	"github.com/dxe/service/data"
)

var (
	Port            = getEnvWithFallback("PORT", "3333")
	Dsn             = getEnvWithFallback("DB_CONNECTION_STRING", "adb_user:adbpassword@tcp(localhost:3306)/campaign_mailer")
	RecaptchaSecret = getEnvWithFallback("RECAPTCHA_SECRET", "")
)

func getEnvWithFallback(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

var helpTheChickensRecipients = []string{
	"kmcdonnell@cityofpetaluma.org",
	"knau@cityofpetaluma.org",
	"jshribbs@cityofpetaluma.org",
	"fquint@cityofpetaluma.org",
	"Jcaderthompson@cityofpetaluma.org",
	"adecarli@cityofpetaluma.org",
	"bbarnacle@cityofpetaluma.org"}

type EmailSettings struct {
	FromDomain string
	Subject    string
	To         func(city data.Municipality, zip data.Zip) []string
}

var CampaignEmailSettings = map[string]EmailSettings{
	"duck":    {FromDomain: "helptheducks.com", Subject: "Prosecute Reichardt Duck Farm for Animal Abuse", To: StaticRecipientList("carla.rodriguez@sonoma-county.org")},
	"chicken": {FromDomain: "helpthechickens.com", Subject: "Shut Down Perdue's Petaluma Poultry Slaughterhouse to End Animal Abuse and Protect Public Health", To: StaticRecipientList(helpTheChickensRecipients...)},
	"sonoma":  {FromDomain: "righttorescue.com", Subject: "Prosecute animal cruelty, not animal rescuers", To: StaticRecipientList("carla.rodriguez@sonoma-county.org")},
	"ridglan": {FromDomain: "righttorescue.com", Subject: "Prosecute animal abuse at Ridglan Farms", To: StaticRecipientList("ismael.ozanne@da.wi.gov")},
	"freezoe": {FromDomain: "freezoe.org", Subject: "Pardon Zoe Rosenberg", To: StaticRecipientList("gavin.newsom@gov.ca.gov")},
	"test":    {FromDomain: "righttorescue.com", Subject: "Test", To: StaticRecipientList("tech@directactioneverywhere.com")},
}

func StaticRecipientList(to ...string) func(city data.Municipality, zip data.Zip) []string {
	return func(city data.Municipality, zip data.Zip) []string {
		return to
	}
}
