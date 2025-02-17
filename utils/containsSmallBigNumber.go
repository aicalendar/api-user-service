package utils

import (
	"unicode"
)

func ContainsSmallBigNumber(string string) bool {
	hasUpper := false
	hasLower := false
	hasDigit := false

	// Loop through each character in the password
	for _, char := range string {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}

	// Return true if all conditions are met
	return hasUpper && hasLower && hasDigit
}
