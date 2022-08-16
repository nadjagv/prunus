package kontroler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	//model "korisnici-servis/model"
	"korisnici-servis/model"
	servis "korisnici-servis/servis"
	util "korisnici-servis/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("prunus_jwt_kljuc")

func OtkrijEndpointe() {
	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {

		var payload util.Kredencijali
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		korisnik, err := servis.ProveriKredencijale(payload)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		rokIsteka := time.Now().Add(time.Minute * 5)
		claims := &util.Claims{
			Email: korisnik.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: rokIsteka.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenStr, err := token.SignedString(jwtKey)
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = tokenStr
		cookie.Expires = rokIsteka

		c.Cookie(cookie)

		korisnikTokenInfo := &model.KorisnikTokenInfo{
			Email: korisnik.Email,
			Tip:   korisnik.Tip,
			Token: tokenStr,
		}
		return c.Status(fiber.StatusOK).JSON(korisnikTokenInfo)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		korisnici := servis.PreuzmiSve()
		return c.Status(fiber.StatusOK).JSON(korisnici)
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		korisnik, err := servis.PreuzmiPoId(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(korisnik)
	})

	app.Get("/email/:email", func(c *fiber.Ctx) error {
		email := c.Params("email")
		korisnik, err := servis.PreuzmiPoEmail(email)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(korisnik)
	})

	log.Fatal(app.Listen(":8082"))

}
