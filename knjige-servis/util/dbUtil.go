package util

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "knjige-servis/model"
)

var knjige = []model.Knjiga{
	{Isbn: "9788677106294",
		Naziv:            "Gordost i predrasuda",
		ImeAutora:        "Džejn",
		PrezimeAutora:    "Ostin",
		Opis:             "Gordost i predrasuda, nezaboravni klasik Džejn Ostin, donosi priču o slobodoumnoj devojci Elizabet Benet koja, kao i njene četiri sestre, mora da se uda za bogatog muža. Suočavajući se sa arogantim, imućnim gospodinom Darsijem, Elizabet se upušta u pronicljiva razmišljanja o životu, porodici i tradiciji, u jednom od najlepših ljubavnih romana svih vremena.",
		Zanr:             model.ZanrEnum(1),
		BrojStrana:       367,
		GodinaNastanka:   time.Date(1813, time.January, 28, 0, 0, 0, 0, time.Local),
		UkupnaKolicina:   10,
		TrenutnoDostupno: 10,
		ProsecnaOcena:    5.0,
		Slika:            "9788677106294.jpg"},
}

var Database *gorm.DB
var err error

func KonektujPopuniDB() {

	dsn := "host=localhost user=postgres password=admin dbname=prunus-knjige-servis-db port=5432 sslmode=disable"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Konektovan na bazu.")
	}

	Database.Migrator().DropTable("knjige")

	Database.AutoMigrate(&model.Knjiga{})

	for _, knjiga := range knjige {
		Database.Create(&knjiga)
	}

}
