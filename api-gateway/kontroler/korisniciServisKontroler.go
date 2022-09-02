package kontroler

import (
	"bytes"
	"fmt"

	"net/http"

	dto "api-gateway/dto"
	util "api-gateway/util"

	"github.com/gofiber/fiber/v2"
)

func RutirajKorisniciServis(app *fiber.App) {
	var prefiks = "/korisnici"
	var korisniciServisUrl = "http://localhost:8082/"

	app.Post(prefiks+"/login", func(c *fiber.Ctx) error {
		response, err := http.Post(korisniciServisUrl+"login", "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.KorisnikTokenInfo
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks, func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Get(korisniciServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.KorisnikDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, _, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		response, err := http.Get(korisniciServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.KorisnikDTO
		err = util.GetJson(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/email/:email", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		mail, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + mail + "\n")
		email := c.Params("email")
		response, err := http.Get(korisniciServisUrl + "email/" + email)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.KorisnikDTO
		err = util.GetJson(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiks, func(c *fiber.Ctx) error {
		response, err := http.Post(korisniciServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Put(prefiks, func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, _, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		request, err := http.NewRequest(http.MethodPut, korisniciServisUrl, bytes.NewBuffer(c.Body()))
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
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodDelete, korisniciServisUrl+idStr, nil)
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

	app.Put(prefiks+"/sumnjiv/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, korisniciServisUrl+"sumnjiv/"+idStr, bytes.NewBuffer(c.Body()))
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

	app.Put(prefiks+"/lozinka", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, _, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		request, err := http.NewRequest(http.MethodPut, korisniciServisUrl+"lozinka", bytes.NewBuffer(c.Body()))
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

	app.Put(prefiks+"/blokiraj/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, korisniciServisUrl+"blokiraj/"+idStr, bytes.NewBuffer(c.Body()))
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

	app.Put(prefiks+"/produzi-clanarinu/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		request, err := http.NewRequest(http.MethodPut, korisniciServisUrl+"produzi-clanarinu/"+idStr, bytes.NewBuffer(c.Body()))
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

}
