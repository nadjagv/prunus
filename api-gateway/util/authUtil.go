package util

import (
	"api-gateway/dto"
	"bytes"

	"github.com/golang-jwt/jwt"

	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

type DozvolaPristupa struct {
	Clan        bool
	Bibliotekar bool
	Admin       bool
}

var jwtKey = []byte("prunus_jwt_kljuc")
var korisniciServisAuthUrl = "http://localhost:8082/auth"

func Auth(c *fiber.Ctx, dozvola DozvolaPristupa) bool {
	authHeaderStr := string(c.Request().Header.Peek("Authorization"))
	if authHeaderStr == "" {
		return false
	}

	body, _ := json.Marshal(dozvola)
	request, err := http.NewRequest(http.MethodPost, korisniciServisAuthUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return false
	}
	request.Header.Set("Authorization", authHeaderStr)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if response.StatusCode == fiber.StatusOK {
		return true
	}
	return false
}
