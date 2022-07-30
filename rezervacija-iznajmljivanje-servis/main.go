package main

import (
	"log"

	util "rezervacija-iznajmljivanje-servis/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	util.KonektujPopuniDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
