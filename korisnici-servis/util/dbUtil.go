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
	{Email: "admin",
		Lozinka:        string(hashLozinka1),
		Ime:            "Admin",
		Prezime:        "Sistemić",
		Tip:            model.TipKorisnika(2),
		IstekClanarine: time.Time{},
		Sumnjiv:        false,
		Blokiran:       false,
	},
	{Email: "bibliotekar",
		Lozinka:        string(hashLozinka1),
		Ime:            "Ljubica",
		Prezime:        "Knjižić",
		Tip:            model.TipKorisnika(1),
		IstekClanarine: time.Time{},
		Sumnjiv:        false,
		Blokiran:       false,
	},
	{Email: "nadjaimatijaprojekat@gmail.com",
		Lozinka:        string(hashLozinka1),
		Ime:            "Pera",
		Prezime:        "Perić",
		Tip:            model.TipKorisnika(0),
		IstekClanarine: time.Time{},
		Sumnjiv:        true,
		Blokiran:       false,
	},
	{Email: "gvozdenacn@gmail.com",
		Lozinka:        string(hashLozinka1),
		Ime:            "Nađa",
		Prezime:        "Gvozdenac",
		Tip:            model.TipKorisnika(0),
		IstekClanarine: time.Time{},
		Sumnjiv:        false,
		Blokiran:       false,
	},
	{Email: "projekatxml@gmail.com",
		Lozinka:        string(hashLozinka1),
		Ime:            "Jovanka",
		Prezime:        "Tomić",
		Tip:            model.TipKorisnika(0),
		IstekClanarine: time.Time{},
		Sumnjiv:        true,
		Blokiran:       true,
	},
}

var Database *gorm.DB
var errDb error

func KonektujPopuniDB() {

	dsn := "host=localhost user=postgres password=admin dbname=prunus-korisnici-servis-db port=5432 sslmode=disable"
	Database, errDb = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDb != nil {
		log.Fatal(errDb)
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
