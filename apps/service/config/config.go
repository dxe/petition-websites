package config

import (
	"os"

	"github.com/dxe/service/data"
)

var (
	Port                      = getEnvWithFallback("PORT", "3333")
	Dsn                       = getEnvWithFallback("DB_CONNECTION_STRING", "adb_user:adbpassword@tcp(localhost:3306)/campaign_mailer")
	RecaptchaSecret           = getEnvWithFallback("RECAPTCHA_SECRET", "")
	GoogleMapsGeocodingAPIKey = getEnvWithFallback("GOOGLE_MAPS_GEOCODING_API_KEY", "")
)

func getEnvWithFallback(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func GetGoogleMapsGeocodingAPIKey() string {
	return GoogleMapsGeocodingAPIKey
}

var helpTheChickensRecipients = []string{
	"kmcdonnell@cityofpetaluma.org",
	"knau@cityofpetaluma.org",
	"jshribbs@cityofpetaluma.org",
	"fquint@cityofpetaluma.org",
	"Jcaderthompson@cityofpetaluma.org",
	"adecarli@cityofpetaluma.org",
	"bbarnacle@cityofpetaluma.org",
}

var factoryFarmWatchRecipients = []string{
	"gavin.newsom@gov.ca.gov",
	"rob.bonta@doj.ca.gov",
	"Senator.Limon@senate.ca.gov",
	"assemblymember.rivas@asm.ca.gov",
}

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
	"freezoe": {FromDomain: "freezoe.org", Subject: "Pardon Zoe Rosenberg", To: StaticRecipientList("gavin.newsom@gov.ca.gov", "pardons@gov.ca.gov")},
	"factoryfarmwatch": {FromDomain: "petition.factoryfarmwatch.org", Subject: "Regulate Factory Farms", To: func(city data.Municipality, zip data.Zip) []string {
		return mergeSliceWtihDeduplication(factoryFarmWatchRecipients, getEmailsForAssemblyMembers(GetAssemblyMembers(city, zip)))
	}},
	"test": {FromDomain: "righttorescue.com", Subject: "Test", To: StaticRecipientList("tech@directactioneverywhere.com")},
}

func StaticRecipientList(to ...string) func(city data.Municipality, zip data.Zip) []string {
	return func(city data.Municipality, zip data.Zip) []string {
		return to
	}
}

func mergeSliceWtihDeduplication[T comparable](slice1 []T, slice2 []T) []T {
	mergedSlice := append(slice1, slice2...)
	seen := make(map[T]bool)
	result := []T{}
	for _, item := range mergedSlice {
		if _, exists := seen[item]; !exists {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Returns all possible assembly members for a given city and zip code.
// Returns only assembly members that match to both the city and zip code.
func GetAssemblyMembers(city data.Municipality, zip data.Zip) map[data.District]data.AssemblyMember {
	potentialDistricts := make(map[data.District]bool)

	assemblyMembersForCity := data.MunicipalityDistrictPercent[city]
	assemblyMembersForZip := data.ZipDistrictPercent[zip]

	// Find districts that are present in both city and zip code mappings
	// where a nonzero percentage of the population lives in that district
	for district, percent := range assemblyMembersForZip {
		if percent > 0 && assemblyMembersForCity[district] > 0 {
			potentialDistricts[district] = true
		}
	}

	members := make(map[data.District]data.AssemblyMember)
	for district := range potentialDistricts {
		member, ok := data.AssemblyMemberData[district]
		if ok {
			members[district] = member
		}
	}
	return members
}

func getEmailsForAssemblyMembers(members map[data.District]data.AssemblyMember) []string {
	emails := make([]string, 0, len(members))
	for _, member := range members {
		emails = append(emails, member.Email)
	}
	return emails
}
