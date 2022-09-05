package kontroler

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	dto "api-gateway/dto"
	util "api-gateway/util"

	"github.com/gofiber/fiber/v2"
)

func RutirajRezIznServis(app *fiber.App) {
	var prefiksRez = "/rezervacije"
	var prefiksIzn = "/iznajmljivanja"
	var iznajmljivanjeServisUrl = "http://localhost:8083/iznajmljivanja/"
	var rezervacijaServisUrl = "http://localhost:8083/rezervacije/"
	var knjigeServisUrl = "http://localhost:8081/"

	//rezervacije
	app.Get(prefiksRez, func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Get(rezervacijaServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.RezervacijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var rezultat []dto.RezervacijaNazivKnjigeDTO
		var knjiga dto.KnjigaSlikaDTO
		for _, rez := range body {
			responseKnjige, err2 := http.Get(knjigeServisUrl + strconv.FormatUint(uint64(rez.KnjigaId), 10))
			if err2 != nil {
				return c.Status(fiber.ErrBadRequest.Code).JSON(err2)
			}

			err = util.GetJson(responseKnjige, &knjiga)
			if err != nil {
				return c.Status(fiber.ErrBadRequest.Code).JSON(err)
			}

			rezDto := dto.RezervacijaNazivKnjigeDTO{
				Id:               rez.Id,
				DatumVremeIsteka: rez.DatumVremeIsteka,
				KorisnikId:       rez.KorisnikId,
				KnjigaId:         rez.KnjigaId,
				KnjigaNaziv:      knjiga.Naziv,
				Aktivno:          rez.Aktivno,
			}
			rezultat = append(rezultat, rezDto)
		}

		return c.Status(response.StatusCode).JSON(rezultat)
	})

	app.Get(prefiksRez+"/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 && tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		response, err := http.Get(rezervacijaServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.RezervacijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		responseKnjige, err2 := http.Get(knjigeServisUrl + strconv.FormatUint(uint64(body.KnjigaId), 10))
		if err2 != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err2)
		}

		var knjiga dto.KnjigaSlikaDTO
		err = util.GetJson(responseKnjige, &knjiga)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		rezultat := dto.RezervacijaNazivKnjigeDTO{
			Id:               body.Id,
			DatumVremeIsteka: body.DatumVremeIsteka,
			KorisnikId:       body.KorisnikId,
			KnjigaId:         body.KnjigaId,
			KnjigaNaziv:      knjiga.Naziv,
			Aktivno:          body.Aktivno,
		}
		return c.Status(response.StatusCode).JSON(rezultat)
	})

	app.Get(prefiksRez+"/korisnik/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		response, err := http.Get(rezervacijaServisUrl + "korisnik/" + idStr)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.RezervacijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var rezultat []dto.RezervacijaNazivKnjigeDTO
		var knjiga dto.KnjigaSlikaDTO
		for _, rez := range body {
			responseKnjige, err2 := http.Get(knjigeServisUrl + strconv.FormatUint(uint64(rez.KnjigaId), 10))
			if err2 != nil {
				fmt.Println(err2)
				return c.Status(fiber.ErrBadRequest.Code).JSON(err2)
			}

			err = util.GetJson(responseKnjige, &knjiga)
			if err != nil {
				fmt.Println(err)
				return c.Status(fiber.ErrBadRequest.Code).JSON(err)
			}

			rezDto := dto.RezervacijaNazivKnjigeDTO{
				Id:               rez.Id,
				DatumVremeIsteka: rez.DatumVremeIsteka,
				KorisnikId:       rez.KorisnikId,
				KnjigaId:         rez.KnjigaId,
				KnjigaNaziv:      knjiga.Naziv,
				Aktivno:          rez.Aktivno,
			}
			rezultat = append(rezultat, rezDto)
		}

		return c.Status(response.StatusCode).JSON(rezultat)
	})

	app.Get(prefiksRez+"/knjiga-korisnik/:knjigaId/:korisnikId", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		knjigaId := c.Params("knjigaId")
		korisnikId := c.Params("korisnikId")
		response, err := http.Get(rezervacijaServisUrl + "knjiga-korisnik/" + knjigaId + "/" + korisnikId)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.RezervacijaDTO
		err = util.GetJson(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiksRez, func(c *fiber.Ctx) error {
		fmt.Println("ejjejefgws")
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 && tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Post(rezervacijaServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Delete(prefiksRez+"/:id", func(c *fiber.Ctx) error {
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
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 && tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
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
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Get(iznajmljivanjeServisUrl)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.IznajmljivanjeDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiksIzn+"/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 && tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		response, err := http.Get(iznajmljivanjeServisUrl + idStr)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body dto.IznajmljivanjeDTO
		err = util.GetJson(response, &body)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Get(prefiksIzn+"/aktivna-korisnik/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 && tip != 1 && tip != 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		idStr := c.Params("id")
		response, err := http.Get(iznajmljivanjeServisUrl + "aktivna-korisnik/:" + idStr)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}

		var body []dto.IznajmljivanjeDTO
		err = util.GetJson(response, &body)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(response.StatusCode).JSON(body)
	})

	app.Post(prefiksIzn, func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Post(iznajmljivanjeServisUrl, "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Post(prefiksIzn+"/vrati", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
		response, err := http.Post(iznajmljivanjeServisUrl+"vrati", "application/json", bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.SendStatus(response.StatusCode)
	})

	app.Put(prefiksIzn+"/produzi/:id", func(c *fiber.Ctx) error {
		authHeaderStr := string(c.Request().Header.Peek("Authorization"))
		email, tip, err := util.Autentifikuj(authHeaderStr[7:])
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if tip != 0 && tip != 1 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		print("Zahtev poslao: " + email + "\n")
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
