package util

import (
	"errors"
	model "korisnici-servis/model"

	"github.com/golang-jwt/jwt"
)

type Kredencijali struct {
	Email   string
	Lozinka string
}

type Claims struct {
	Email string
	Tip   model.TipKorisnika
	jwt.StandardClaims
}

type DozvolaPristupa struct {
	Clan        bool
	Bibliotekar bool
	Admin       bool
}

var jwtKey = []byte("prunus_jwt_kljuc")

func Autentifikuj(tokenStr string) (string, model.TipKorisnika, error) {
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

func Autorizuj(dozvola DozvolaPristupa, tipZaProveru model.TipKorisnika) bool {
	if tipZaProveru == model.CLAN && dozvola.Clan {
		return true
	} else if tipZaProveru == model.BIBLIOTEKAR && dozvola.Bibliotekar {
		return true
	} else if tipZaProveru == model.ADMIN && dozvola.Admin {
		return true
	}
	return false
}
