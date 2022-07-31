package kontroler

import (
	"log"

	servis "knjige-servis/servis"

	"github.com/gofiber/fiber/v2"
)

func OtkrijEndpointe() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		knjige := servis.PreuzmiSve()
		return c.JSON(knjige)
	})

	log.Fatal(app.Listen(":8081"))

}
