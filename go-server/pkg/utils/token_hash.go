package utils

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func TokensEqualHash(storedHash string, token string) bool {
	incomingHash := HashToken(token)
	return subtle.ConstantTimeCompare([]byte(storedHash), []byte(incomingHash)) == 1
}
