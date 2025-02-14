package passwords

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

func ComparePasswords(inputPassword, storedHash, storedSalt string) bool {

	// Decode the base64 encoded salt and hash
	salt, _ := base64.StdEncoding.DecodeString(storedSalt)

	// Argon2id parameters (must match the hashing function)
	time := uint32(3)
	memory := uint32(64 * 1024)
	threads := uint8(1)
	keyLen := uint32(32)

	// Hash input password with the same salt
	inputHash := argon2.IDKey([]byte(inputPassword), salt, time, memory, threads, keyLen)

	// Compare the newly generated hash with the stored hash
	return base64.StdEncoding.EncodeToString(inputHash) == storedHash
}
