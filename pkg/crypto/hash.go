package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateSaltedHash ...
func GenerateSaltedHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// CompareHashAndPassword ...
func CompareHashAndPassword(existingHash string, inc string) error {
	incoming := []byte(inc)
	existing := []byte(existingHash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
