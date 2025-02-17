package utils

import (
	"unicode/utf8"
)

func HasEnoughLetters(string string, count int) bool {
	length := utf8.RuneCountInString(string)

	return length >= count
}
