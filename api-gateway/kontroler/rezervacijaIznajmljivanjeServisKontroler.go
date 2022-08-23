package kontroler

import (
	"bytes"
	"fmt"
	"net/http"

	dto "api-gateway/dto"
	util "api-gateway/util"

	"github.com/gofiber/fiber/v2"
)

func RutirajRezIznServis(app *fiber.App) {
	var prefiksRez = "/rezervacije"
	var prefiksIzn = "/iznajmljivanja"
	var iznajmljivanjeServisUrl = "http://localhost:8081/iznajmljivanja/"
	var rezervacijaServisUrl = "http://localhost:8081/rezervacije/"

	//rezervacije
	app.Get(prefiksRez, func(c *fiber.Ctx) error {
		response, err := http.Get(rezervacijaServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.RezervacijaDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiksRez+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(rezervacijaServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.RezervacijaDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiksRez, func(c *fiber.Ctx) error {
		response, err := http.Post(rezervacijaServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Delete(prefiksRez+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodDelete, rezervacijaServisUrl+idStr, nil)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		return c.SendStatus(response.StatusCode)
	})

	app.Put(prefiksRez+"/otkazi/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, rezervacijaServisUrl+"otkazi/"+idStr, bytes.NewBuffer(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	//iznajmljivanja
	app.Get(prefiksIzn, func(c *fiber.Ctx) error {
		response, err := http.Get(iznajmljivanjeServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.IznajmljivanjeDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiksIzn+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(iznajmljivanjeServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.IznajmljivanjeDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiksIzn+"/aktivna-korisnik/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(iznajmljivanjeServisUrl + "aktivna-korisnik/:" + idStr)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.IznajmljivanjeDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiksIzn, func(c *fiber.Ctx) error {
		response, err := http.Post(iznajmljivanjeServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Post(prefiksIzn+"/vrati", func(c *fiber.Ctx) error {
		response, err := http.Post(iznajmljivanjeServisUrl+"vrati", "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Put(prefiksIzn+"/produzi/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, iznajmljivanjeServisUrl+"produzi/"+idStr, bytes.NewBuffer(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Delete(prefiksIzn+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodDelete, iznajmljivanjeServisUrl+idStr, nil)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		return c.SendStatus(response.StatusCode)
	})

}
