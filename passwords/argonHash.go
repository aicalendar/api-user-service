package passwords

import (
	"encoding/base64"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/argon2"
)

// Hash a password using Argon2id with a unique salt
func HashPassword(password string) (string, string, error) {
	// Generate random salt
	salt, err := GenerateSalt(16)
	if err != nil {
		log.Error().Msg("Failed generating salt")
		return "", "", nil
	}

	// Argon2id parameters
	time := uint32(3)
	memory := uint32(64 * 1024)
	threads := uint8(1)
	keyLen := uint32(32)

	// Generate Argon2id hash
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Encode salt and hash as base64 for storage
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	encodedHash := base64.StdEncoding.EncodeToString(hash)

	return encodedHash, encodedSalt, nil
}
