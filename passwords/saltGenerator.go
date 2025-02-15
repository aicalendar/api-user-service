package passwords

import (
	"crypto/rand"
)

// Generate a random salt
func GenerateSalt(size int) ([]byte, error) {

	// Generate random bytes
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}
