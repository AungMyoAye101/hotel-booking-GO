package utils

import (
	"crypto/subtle"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyPassword returns (ok, upgradedHash, err).
// If the stored password is plaintext (legacy), upgradedHash will contain a bcrypt hash when ok=true.
func VerifyPassword(stored string, plain string) (bool, string, error) {
	// bcrypt hashes start with "$2"
	if strings.HasPrefix(stored, "$2") {
		if err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(plain)); err != nil {
			return false, "", nil
		}
		return true, "", nil
	}

	if subtle.ConstantTimeCompare([]byte(stored), []byte(plain)) != 1 {
		return false, "", nil
	}
	upgraded, err := HashPassword(plain)
	if err != nil {
		return true, "", err
	}
	return true, upgraded, nil
}
