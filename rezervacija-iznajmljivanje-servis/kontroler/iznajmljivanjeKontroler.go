package kontroler

import (
	"strconv"

	model "rezervacija-iznajmljivanje-servis/model"
	servis "rezervacija-iznajmljivanje-servis/servis"

	"github.com/gofiber/fiber/v2"
)

func OtkrijEndpointeIzn(app *fiber.App) {
	prefiks := "/iznajmljivanja"

	app.Get(prefiks+"/", func(c *fiber.Ctx) error {
		rez := servis.PreuzmiSveIzn()
		var rezultat []model.IznajmljivanjeDTO
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
		rez, err := servis.PreuzmiPoIdIzn(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(rez.MapirajNaDTO())
	})

	app.Get(prefiks+"/knjiga-korisnik/:knjigaId/:korisnikId", func(c *fiber.Ctx) error {
		idStr := c.Params("korisnikId")
		korisnikId, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		idStr = c.Params("knjigaId")
		knjigaId, err2 := strconv.ParseUint(idStr, 10, 64)
		if err2 != nil {
			return fiber.NewError(fiber.StatusBadRequest, err2.Error())
		}

		rez := servis.PreuzmiAktivnuKorisnikKnjigaIzn(uint(korisnikId), uint(knjigaId))
		return c.Status(fiber.StatusOK).JSON(rez.MapirajNaDTO())
	})

	app.Get(prefiks+"/aktivna-korisnik/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		rez := servis.PreuzmiPoKorisnikuAktivnaIzn(uint(id))
		var rezultat []model.IznajmljivanjeDTO
		for _, r := range rez {
			rezultat = append(rezultat, r.MapirajNaDTO())
		}
		return c.Status(fiber.StatusOK).JSON(rezultat)
	})

	app.Post(prefiks+"/", func(c *fiber.Ctx) error {
		var payload model.IznajmljivanjeDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.KreirajIzn(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post(prefiks+"/vrati", func(c *fiber.Ctx) error {
		var payload model.IznajmljivanjeDTO
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.VratiKnjigu(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put(prefiks+"/produzi/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.ProduziIzn(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete(prefiks+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.ObrisiPoIdIzn(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

}
