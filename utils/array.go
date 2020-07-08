package util

import "strings"

// Find the value in the array.
func Find(arr []string, value string) string {
	for _, v := range arr {
		if v == value {
			return value
		}
	}

	return ""
}

// GetCountry name.
func GetCountry(country []string) string {
	var matchedCountry string
	if len(country) > 1 {
		matchedCountry = strings.Join(country, "-")
	} else {
		matchedCountry = country[0]
	}
	matchedCountry = strings.ToLower(matchedCountry)

	return matchedCountry
}
