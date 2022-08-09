package util

import (
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
