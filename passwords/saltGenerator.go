package passwords

import (
	"crypto/rand"
	"github.com/rs/zerolog/log"
)

// Generate a random salt
func GenerateSalt(size int) []byte {

	// Generate random bytes
	salt := make([]byte, size)
	_, err := rand.Read(salt)

	// Log on error
	if err != nil {
		log.Error().Msg("Failed reading random salt into bytes!")
		return nil
	}

	return salt
}
