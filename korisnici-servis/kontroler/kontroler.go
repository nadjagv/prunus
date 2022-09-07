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

		rokIsteka := time.Now().Add(time.Minute * 30)
		claims := &util.Claims{
			Email: korisnik.Email,
			Tip:   korisnik.Tip,
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
			Id:    korisnik.ID,
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

	app.Post("/", func(c *fiber.Ctx) error {
		var payload model.KorisnikDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Kreiraj(payload)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/", func(c *fiber.Ctx) error {
		var payload model.KorisnikDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Izmeni(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/lozinka", func(c *fiber.Ctx) error {
		var payload model.IzmenaLozinkeDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.IzmeniLozinku(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/sumnjiv/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.OznaciSumnjiv(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.ObrisiPoId(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/produzi-clanarinu/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.ProduziClanarinu(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/opomeni/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Opomeni(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/blokiraj/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Println("ojj")
			fmt.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Blokiraj(uint(id), string(c.Body()))
		if err != nil {
			fmt.Println("hrh")
			fmt.Println(err)
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/odblokiraj/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Odblokiraj(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/pretrazi/:param", func(c *fiber.Ctx) error {
		param := c.Params("param")
		korisnici := servis.Pretrazi(param)
		return c.Status(fiber.StatusOK).JSON(korisnici)
	})

	log.Fatal(app.Listen(":8082"))

}
