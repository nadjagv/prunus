package kontroler

import (
	"bytes"
	"fmt"
	"net/http"

	dto "api-gateway/dto"
	util "api-gateway/util"

	"github.com/gofiber/fiber/v2"
)

func RutirajKnjigeServis(app *fiber.App) {
	var prefiks = "/knjige"
	var knjigeServisUrl = "http://localhost:8081/"
	app.Get(prefiks, func(c *fiber.Ctx) error {
		// authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		// email, err := util.Autentifikuj(authHeaderStr[7:])
		// if err != nil {
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }
		// print("Zahtev poslao: " + email + "\n")
		response, err := http.Get(knjigeServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.KnjigaSlikaDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(knjigeServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.KnjigaSlikaDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/isbn/:isbn", func(c *fiber.Ctx) error {
		isbn := c.Params("isbn")
		response, err := http.Get(knjigeServisUrl + "isbn/" + isbn)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.KnjigaSlikaDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiks, func(c *fiber.Ctx) error {
		response, err := http.Post(knjigeServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Put(prefiks, func(c *fiber.Ctx) error {
		request, err := http.NewRequest(http.MethodPut, knjigeServisUrl, bytes.NewBuffer(c.Body()))
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

	app.Delete(prefiks+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodDelete, knjigeServisUrl+idStr, nil)
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

	app.Get(prefiks+"/kolicina/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		response, err := http.Get(knjigeServisUrl + "kolicina/" + id)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body uint
		err = util.GetJsonIC(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Put(prefiks+"/smanji-kolicinu/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"smanji-kolicinu/"+idStr, bytes.NewBuffer(c.Body()))
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

	app.Put(prefiks+"/povecaj-kolicinu/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"povecaj-kolicinu/"+idStr, bytes.NewBuffer(c.Body()))
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

	app.Get(prefiks+"/pretplata/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		response, err := http.Get(knjigeServisUrl + "pretplata/" + id)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.PretplataDTO
		err = util.GetJsonIC(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiks+"/pretplata", func(c *fiber.Ctx) error {
		response, err := http.Post(knjigeServisUrl+"pretplata/", "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Delete(prefiks+"/pretplata/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodDelete, knjigeServisUrl+"pretplata/"+idStr, nil)
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
