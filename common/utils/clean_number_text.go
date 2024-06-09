package utils

import "regexp"

func CleanNumberText(data string) string {
	re := regexp.MustCompile("[^0-9]+")
	return re.ReplaceAllString(data, "")
}
