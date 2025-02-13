package passwords

import (
	"crypto/rand"
)

// Generate a random salt
func GenerateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt) // Secure random bytes
	if err != nil {
		return nil, err
	}
	return salt, nil
}
