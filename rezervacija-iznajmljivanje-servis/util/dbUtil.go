package util

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "rezervacija-iznajmljivanje-servis/model"
)

var Database *gorm.DB

func KonektujPopuniDB() {

	dsn := "host=localhost user=postgres password=admin dbname=prunus-rezervacija-iznajmljivanje-servis-db port=5432 sslmode=disable"
	Database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Konektovan na bazu.")
	}

	Database.Migrator().DropTable("rezervacije")
	Database.Migrator().DropTable("iznajmljivanja")

	Database.AutoMigrate(&model.Rezervacija{})
	Database.AutoMigrate(&model.Iznajmljivanje{})

}
