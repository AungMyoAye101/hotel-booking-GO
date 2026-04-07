package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type JWTClaims struct {
	Sub  string `json:"sub"`
	Aud  string `json:"aud"`
	Role string `json:"role,omitempty"`
	Typ  string `json:"typ"` // "access" | "refresh"
	Exp  int64  `json:"exp"`
	Iat  int64  `json:"iat"`
	Jti  string `json:"jti,omitempty"`
}

func SignHS256(claims JWTClaims, secret string) (string, error) {
	if secret == "" {
		return "", errors.New("secret is required")
	}

	headerJSON, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		return "", err
	}
	payloadJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	enc := base64.RawURLEncoding
	headerPart := enc.EncodeToString(headerJSON)
	payloadPart := enc.EncodeToString(payloadJSON)
	signingInput := headerPart + "." + payloadPart

	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(signingInput))
	sigPart := enc.EncodeToString(mac.Sum(nil))

	return signingInput + "." + sigPart, nil
}

func ParseAndVerifyHS256(token string, secret string) (JWTClaims, error) {
	var claims JWTClaims

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return claims, ErrInvalidToken
	}

	enc := base64.RawURLEncoding

	headerBytes, err := enc.DecodeString(parts[0])
	if err != nil {
		return claims, ErrInvalidToken
	}
	var header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return claims, ErrInvalidToken
	}
	if header.Alg != "HS256" {
		return claims, ErrInvalidToken
	}

	signingInput := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(signingInput))
	expectedSig := mac.Sum(nil)

	sigBytes, err := enc.DecodeString(parts[2])
	if err != nil {
		return claims, ErrInvalidToken
	}
	if !hmac.Equal(sigBytes, expectedSig) {
		return claims, ErrInvalidToken
	}

	payloadBytes, err := enc.DecodeString(parts[1])
	if err != nil {
		return claims, ErrInvalidToken
	}
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return claims, ErrInvalidToken
	}

	if claims.Exp == 0 || claims.Sub == "" || claims.Typ == "" || claims.Aud == "" {
		return claims, ErrInvalidToken
	}

	now := time.Now().Unix()
	if now > claims.Exp {
		return claims, ErrExpiredToken
	}

	return claims, nil
}

func NewAccessClaims(sub, aud, role string, ttl time.Duration) JWTClaims {
	now := time.Now()
	c := JWTClaims{
		Sub: sub,
		Aud: aud,
		Typ: "access",
		Iat: now.Unix(),
		Exp: now.Add(ttl).Unix(),
	}
	if role != "" {
		c.Role = role
	}
	return c
}

func NewRefreshClaims(sub, aud, role, jti string, ttl time.Duration) JWTClaims {
	now := time.Now()
	c := JWTClaims{
		Sub: sub,
		Aud: aud,
		Typ: "refresh",
		Jti: jti,
		Iat: now.Unix(),
		Exp: now.Add(ttl).Unix(),
	}
	if role != "" {
		c.Role = role
	}
	return c
}
