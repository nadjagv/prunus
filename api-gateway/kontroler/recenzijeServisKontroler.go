package kontroler

import (
	// "bytes"
	// "fmt"

	// "net/http"

	// dto "api-gateway/dto"
	// util "api-gateway/util"

	"api-gateway/dto"
	"api-gateway/util"
	"fmt"
	"net/http"

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
		return c.Status(response.StatusCode).JSON(body)
	})

}
