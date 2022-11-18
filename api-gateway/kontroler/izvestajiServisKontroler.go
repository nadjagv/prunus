package kontroler

import (
	"api-gateway/dto"
	"api-gateway/util"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RutirajIzvestajiServis(app *fiber.App) {
	var prefiks = "/izvestaji"
	var izvestajiServisUrl = "http://localhost:8001/"

	app.Get(prefiks, func(c *fiber.Ctx) error {
		dozvola := util.DozvolaPristupa{
			Clan:        false,
			Bibliotekar: true,
			Admin:       true,
		}
		if !util.Auth(c, dozvola) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		pocetak := c.Query("pocetak")
		kraj := c.Query("kraj")

		response, err := http.Get(izvestajiServisUrl + "?pocetak=" + pocetak + "&kraj=" + kraj)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.IzvestajDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

}
