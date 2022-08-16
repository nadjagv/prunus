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
		err = util.GetJsonIC(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

}
