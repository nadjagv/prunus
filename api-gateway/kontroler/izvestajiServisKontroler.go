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
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 2 && tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")

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
