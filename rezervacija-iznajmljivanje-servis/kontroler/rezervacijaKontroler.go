package kontroler

import (
	"strconv"

	model "rezervacija-iznajmljivanje-servis/model"
	servis "rezervacija-iznajmljivanje-servis/servis"

	"github.com/gofiber/fiber/v2"
)

func OtkrijEndpointeRez(app *fiber.App) {
	prefiks := "/rezervacije"

	app.Get(prefiks+"/", func(c *fiber.Ctx) error {
		rez := servis.PreuzmiSveRez()
		var rezultat []model.RezervacijaDTO
		for _, r := range rez {
			rezultat = append(rezultat, r.MapirajNaDTO())
		}
		return c.Status(fiber.StatusOK).JSON(rezultat)
	})

	app.Get(prefiks+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		rez, err := servis.PreuzmiPoIdRez(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(rez.MapirajNaDTO())
	})

	app.Post(prefiks+"/", func(c *fiber.Ctx) error {
		var payload model.RezervacijaDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.KreirajRez(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete(prefiks+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.ObrisiPoIdRez(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

}
