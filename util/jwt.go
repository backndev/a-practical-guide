package util

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

const secretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,
		ExpiresAt: &jwt.Time{
			Time: time.Now().Add(time.Hour * 24),
		},
	})

	return claims.SignedString([]byte(secretKey))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
