package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	kontroler "api-gateway/kontroler"
)

func main() {
	app := fiber.New()

	kontroler.RutirajKnjigeServis(app)

	log.Fatal(app.Listen(":8080"))
}
