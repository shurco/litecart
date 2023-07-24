package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
)

type Setting struct {
	Secret            string
	SecretExpireHours int
}

// GenerateNewToken func for generate a new Access token.
func GenerateNewToken(secret, id string, expires int64, credentials []string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["expires"] = expires

	for _, credential := range credentials {
		claims[credential] = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
