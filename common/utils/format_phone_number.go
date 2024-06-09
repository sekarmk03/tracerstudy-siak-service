package utils

import (
	"regexp"
	"strings"
)

func FormatPhoneNumber(phoneNumber string) string {
	// Split by '/', ',' and take the first part
	if strings.Contains(phoneNumber, "/") {
		parts := strings.Split(phoneNumber, "/")
		phoneNumber = parts[0]
	} else if strings.Contains(phoneNumber, ",") {
		parts := strings.Split(phoneNumber, ",")
		phoneNumber = parts[0]
	}

	// Remove all non-numeric characters
	re := regexp.MustCompile("[^0-9]+")
	cleanedNumber := re.ReplaceAllString(phoneNumber, "")

	// If the cleaned number starts with "0", replace it with "62"
	if strings.HasPrefix(cleanedNumber, "0") {
		cleanedNumber = "62" + cleanedNumber[1:]
	}
	if strings.HasPrefix(cleanedNumber, "8") || strings.HasPrefix(cleanedNumber, "9") {
		cleanedNumber = "62" + cleanedNumber
	}

	return cleanedNumber
}
