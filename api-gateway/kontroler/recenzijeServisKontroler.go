package kontroler

import (
	"api-gateway/dto"
	"api-gateway/util"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RutirajRecenzijeServis(app *fiber.App) {
	var prefiks = "/recenzije"
	var recenzijeServisUrl = "http://localhost:8000/"

	app.Get(prefiks, func(c *fiber.Ctx) error {
		response, err := http.Get(recenzijeServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		// bodyBytes, err := io.ReadAll(response.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)

		var body []dto.RecenzijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(recenzijeServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.RecenzijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/knjiga/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(recenzijeServisUrl + "knjiga/" + idStr)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.RecenzijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/knjiga-odobreni/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		response, err := http.Get(recenzijeServisUrl + "knjiga-odobreni/" + idStr)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.RecenzijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiks+"/pregled/sve", func(c *fiber.Ctx) error {
		response, err := http.Get(recenzijeServisUrl + "pregled/sve")
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.RecenzijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		var rezultat []dto.RecenzijaNazivEmailDTO
		for _, rec := range body {
			rneDTO, err2 := util.MapirajNaRecenzijeNazivEmailDTO(rec)
			if err2 != nil {
				fmt.Println(err)
				return c.Status(fiber.ErrBadRequest.Code).JSON(err)
			}
			rezultat = append(rezultat, rneDTO)
		}

		return c.Status(response.StatusCode).JSON(rezultat)
	})

	app.Put(prefiks+"/odobri/:id", func(c *fiber.Ctx) error {
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
		request, err := http.NewRequest(http.MethodPut, recenzijeServisUrl+"odobri/"+idStr, bytes.NewBuffer(c.Body()))
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

	app.Put(prefiks+"/odbij/:id", func(c *fiber.Ctx) error {
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
		request, err := http.NewRequest(http.MethodPut, recenzijeServisUrl+"odbij/"+idStr, bytes.NewBuffer(c.Body()))
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

	app.Post(prefiks, func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Post(recenzijeServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Get(prefiks+"/postoji/:korisnikId/:knjigaId", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")

		korisnikId := c.Params("korisnikId")
		knjigaId := c.Params("knjigaId")

		response, err := http.Get(recenzijeServisUrl + "postoji/" + korisnikId + "/" + knjigaId)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		bodyBytes, err2 := io.ReadAll(response.Body)
		if err2 != nil {
			fmt.Println(err2)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err2)
		}
		postoji, err3 := strconv.ParseBool(string(bodyBytes))
		if err3 != nil {
			fmt.Println(err3)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err3)
		}

		return c.Status(response.StatusCode).JSON(postoji)
	})

}
