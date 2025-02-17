package utils

import (
	"unicode"
)

func isAlphanumeric(string string) bool {
	// Loop through each character in the string
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) { // Check if character is neither a letter nor a digit
			return false
		}
	}
	return true
}
