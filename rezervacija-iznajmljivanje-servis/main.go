package main

import (
	"log"

	kontroler "rezervacija-iznajmljivanje-servis/kontroler"
	util "rezervacija-iznajmljivanje-servis/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	util.KonektujPopuniDB()

	app := fiber.New()

	kontroler.OtkrijEndpointeIzn(app)
	kontroler.OtkrijEndpointeRez(app)

	log.Fatal(app.Listen(":8083"))
}
