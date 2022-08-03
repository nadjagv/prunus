package kontroler

import (
	"log"
	"strconv"

	model "knjige-servis/model"
	servis "knjige-servis/servis"

	"github.com/gofiber/fiber/v2"
)

func OtkrijEndpointe() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		knjige := servis.PreuzmiSve()
		return c.Status(fiber.StatusOK).JSON(knjige)
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		knjiga, err := servis.PreuzmiPoId(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(knjiga)
	})

	app.Get("/isbn/:isbn", func(c *fiber.Ctx) error {
		isbn := c.Params("isbn")
		knjiga, err := servis.PreuzmiPoIsbn(isbn)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(knjiga)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		var payload model.KnjigaDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Kreiraj(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(nil)
	})

	app.Put("/", func(c *fiber.Ctx) error {
		var payload model.KnjigaDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.Izmeni(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(nil)
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
		return c.Status(fiber.StatusOK).JSON(nil)
	})

	log.Fatal(app.Listen(":8081"))

}
