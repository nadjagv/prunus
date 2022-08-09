package util

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Kredencijali struct {
	Email   string
	Lozinka string
}

type Claims struct {
	Email string
	jwt.StandardClaims
}

var jwtKey = []byte("prunus_jwt_kljuc")

func Autentifikuj(tokenStr string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return "", errors.New("Unauthorized")
	}

	if !token.Valid {
		return "", errors.New("Unauthorized")
	}

	return claims.Email, nil
}
