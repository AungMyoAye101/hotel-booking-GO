package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Kind  string `json:"kind,omitempty"` // "user" | "admin"
	Role  string `json:"role,omitempty"`
	Typ   string `json:"typ,omitempty"` // "access" | "refresh"
	jwt.RegisteredClaims
}

func GenerateAccessToken(sub string, email string, role string, secret []byte) (string, error) {

	kind := "user"
	if role == "admin" || role == "staff" {
		kind = "admin"
	}
	claims := Claims{
		Sub:   sub,
		Email: email,
		Kind:  kind,
		Role:  role,
		Typ:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)

}
func GenerateRefreshToken(sub string, email string, role string, secret []byte) (string, error) {
	kind := "user"
	if role == "admin" || role == "staff" {
		kind = "admin"
	}
	claims := Claims{
		Sub:   sub,
		Email: email,
		Kind:  kind,
		Role:  role,
		Typ:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)

}

func ParseToken(tokenStr string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure signing method is correct
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
