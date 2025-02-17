package utils

import (
	"strings"
)

func IsEmpty(string string) bool {
	return strings.TrimSpace(string) == ""
}
