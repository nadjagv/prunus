package util

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "korisnici-servis/model"
)

var hashLozinka1, err = bcrypt.GenerateFromPassword([]byte("lozinka1"), bcrypt.DefaultCost)

var korisnici = []model.Korisnik{
	{Email: "gvozdenacn@gmail.com",
		Lozinka:        string(hashLozinka1),
		Ime:            "Nađa",
		Prezime:        "Gvozdenac",
		Tip:            model.TipKorisnika(2),
		IstekClanarine: time.Time{},
		Sumnjiv:        false,
	},
}

var Database *gorm.DB

func KonektujPopuniDB() {

	dsn := "host=localhost user=postgres password=admin dbname=prunus-korisnici-servis-db port=5432 sslmode=disable"
	Database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Konektovan na bazu.")
	}

	Database.Migrator().DropTable("korisnici")

	Database.AutoMigrate(&model.Korisnik{})

	for _, korisnik := range korisnici {
		Database.Create(&korisnik)
	}

	var k model.Korisnik

	Database.First(&k, 1)
	fmt.Println(k.Lozinka)

}
