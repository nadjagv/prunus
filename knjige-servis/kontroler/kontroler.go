package kontroler

import (
	"fmt"
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
		var rezultat []model.KnjigaSlikaDTO
		for _, knjiga := range knjige {
			rezultat = append(rezultat, knjiga.MapirajNaSlikaDTO())
		}
		return c.Status(fiber.StatusOK).JSON(rezultat)
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
		return c.Status(fiber.StatusOK).JSON(knjiga.MapirajNaSlikaDTO())
	})

	app.Get("/isbn/:isbn", func(c *fiber.Ctx) error {
		isbn := c.Params("isbn")
		knjiga, err := servis.PreuzmiPoIsbn(isbn)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(knjiga.MapirajNaSlikaDTO())
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
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/", func(c *fiber.Ctx) error {
		fmt.Println("edit")
		var payload model.KnjigaDTO
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

	app.Get("/kolicina/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		dostupno, err := servis.ProveriDostupnuKolicinu(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(dostupno)
	})

	app.Put("/smanji-kolicinu/:id", func(c *fiber.Ctx) error {
		fmt.Println("smanji")
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.SmanjiDostupnuKolicinu(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/povecaj-kolicinu/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.PovecajDostupnuKolicinu(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	//pretplata
	app.Get("/pretplata/:korisnikId", func(c *fiber.Ctx) error {
		idStr := c.Params("korisnikId")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		pretplate := servis.PreuzmiPoKorisniku(uint(id))

		return c.Status(fiber.StatusOK).JSON(pretplate)
	})

	app.Post("/pretplata", func(c *fiber.Ctx) error {
		var payload model.Pretplata
		err := c.BodyParser(&payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.KreirajPretplatu(payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete("/pretplata/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = servis.ObrisiPoIdPretplatu(uint(id))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":8081"))

}
