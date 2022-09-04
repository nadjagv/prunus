package kontroler

import (
	"log"

	"github.com/gofiber/fiber/v2"

	dto "mejl-servis/dto"
	servis "mejl-servis/servis"
)

func OtkrijEndpointe() {
	app := fiber.New()

	app.Post("/pretplata", func(c *fiber.Ctx) error {
		var payload dto.Mejl
		err := c.BodyParser(&payload)
		if err != nil {
			println(err.Error())
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = servis.PosaljiMejl("Dostupna knjiga", payload.Poruka, payload.MejlAdresa)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/opomena", func(c *fiber.Ctx) error {
		var payload dto.Mejl
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.PosaljiMejl("Opomena", payload.Poruka, payload.MejlAdresa)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/blokiranje", func(c *fiber.Ctx) error {
		var payload dto.Mejl
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.PosaljiMejl("Blokiranje", payload.Poruka, payload.MejlAdresa)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/aktiviranje", func(c *fiber.Ctx) error {
		var payload dto.Mejl
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.PosaljiMejl("Aktiviranje", payload.Poruka, payload.MejlAdresa)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":8084"))

}
