package kontroler

import (
	"fmt"
	"strconv"

	model "rezervacija-iznajmljivanje-servis/model"
	servis "rezervacija-iznajmljivanje-servis/servis"

	"github.com/gofiber/fiber/v2"
)

func OtkrijEndpointeRez(app *fiber.App) {
	prefiks := "/rezervacije"

	servis.ProveravajIstekRezervacija()

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

	app.Get(prefiks+"/korisnik/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		rez := servis.PreuzmiAktivneKorisnikRez(uint(id))
		var rezultat []model.RezervacijaDTO
		for _, r := range rez {
			rezultat = append(rezultat, r.MapirajNaDTO())
		}
		return c.Status(fiber.StatusOK).JSON(rezultat)
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

		rez, err3 := servis.PreuzmiAktivnuKorisnikKnjigaRez(uint(korisnikId), uint(knjigaId))
		if err3 != nil {
			aktivne := servis.PreuzmiAktivneKorisnikRez(uint(korisnikId))
			if len(aktivne) >= 3 {
				return fiber.NewError(fiber.StatusMethodNotAllowed, err3.Error())
			}
		}
		return c.Status(fiber.StatusOK).JSON(rez.MapirajNaDTO())
	})

	app.Post(prefiks+"/", func(c *fiber.Ctx) error {
		var payload model.RezervacijaDTO
		err := c.BodyParser(&payload)
		if err != nil {
			fmt.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.KreirajRez(payload)
		if err != nil {
			fmt.Println(err)
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

	app.Put(prefiks+"/otkazi/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.OtkaziRezervaciju(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

}
