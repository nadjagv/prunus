package util

import (
	"errors"

	"github.com/golang-jwt/jwt"

	dto "api-gateway/dto"
)

type Kredencijali struct {
	Email   string
	Lozinka string
}

type Claims struct {
	Email string
	Tip   dto.TipKorisnika
	jwt.StandardClaims
}

var jwtKey = []byte("prunus_jwt_kljuc")

func Autentifikuj(tokenStr string) (string, dto.TipKorisnika, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return "", 0, errors.New("Unauthorized")
	}

	if !token.Valid {
		return "", 0, errors.New("Unauthorized")
	}

	return claims.Email, claims.Tip, nil
}
