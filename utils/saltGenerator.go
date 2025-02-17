package utils

import (
	"crypto/rand"
	"errors"
)

// Generate a random salt
func GenerateSalt(size int) ([]byte, error) {

	// Handle error size smaller than 1
	if size < 0 {
		return nil, errors.New("Canno't generate salt with byte size smaller than 1!")
	}

	// Generate random bytes
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}
