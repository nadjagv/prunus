package util

import (
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
