package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	kontroler "api-gateway/kontroler"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	kontroler.RutirajKnjigeServis(app)
	kontroler.RutirajKorisniciServis(app)

	log.Fatal(app.Listen(":8080"))
}
