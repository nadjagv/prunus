package util

import (
	"fmt"
	"log"

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
		GodinaNastanka:   1813,
		UkupnaKolicina:   10,
		TrenutnoDostupno: 0,
		ProsecnaOcena:    5.0,
		BrojOcena:        1,
		Slika:            "9788677106294.jpg",
	},
	{Isbn: "9788610014396",
		Naziv:            "Bornova Dominacija",
		ImeAutora:        "Robert",
		PrezimeAutora:    "Ladlam",
		Opis:             "Bornovi neprijatelji postaju sve jači... Vreme je da se uzvrati svom snagom. Severus Domna, drevna tajna organizacija, saziva članove sa četiri kraja sveta kako bi uklonili jedinog čoveka koji može da im poremeti planove – Džejsona Borna. A da bi postigli ono što nikome još nije pošlo za rukom, imaju paklenu ideju: oružje Bornovog uništenja biće jedina osoba kojoj on bezuslovno veruje.",
		Zanr:             model.ZanrEnum(4),
		BrojStrana:       323,
		GodinaNastanka:   2015,
		UkupnaKolicina:   5,
		TrenutnoDostupno: 5,
		ProsecnaOcena:    5.0,
		BrojOcena:        1,
		Slika:            "9788610014396.jpg",
	},
	{Isbn: "9788663030428",
		Naziv:            "Atomske Navike",
		ImeAutora:        "Džejms",
		PrezimeAutora:    "Klir",
		Opis:             "Knjiga pripada žanru psihologije uspeha i od svog prvog objavljivanja pre manje od godinu dana postala je bestseler u toj oblasti. Ona će se ove godine naći na policama u čak 40 zemalja širom sveta! Džejms Klir, jedan od vodećih svetskih stručnjaka za formiranje navika, otkriva praktične strategije koje će nas naučiti kako da formiramo dobre navike, razbijemo loše i ovladamo sitnim navikama koje vode do izvanrednih rezultata.",
		Zanr:             model.ZanrEnum(7),
		BrojStrana:       267,
		GodinaNastanka:   2019,
		UkupnaKolicina:   8,
		TrenutnoDostupno: 8,
		ProsecnaOcena:    5.0,
		BrojOcena:        1,
		Slika:            "9788663030428.jpg",
	},
	{Isbn: "9788677104986",
		Naziv:            "Kako zadobiti prijatelje i uticati na ljude",
		ImeAutora:        "Dejl",
		PrezimeAutora:    "Karnegi",
		Opis:             "Prodata u više od šesnaest miliona primeraka širom sveta, knjiga Kako zadobiti prijatelje i uticati na ljude pruža mnoštvo korisnih saveta za uspostavljanje i održavanje uspešne komunikacije. Možete da tražite posao koji želite – a da ga dobijete! Možete da prihvatite posao kojim se bavite – a da napredujete! Možete da preokrenete bilo koju situaciju – u svoju korist! Karnegijev jasan, već isproban savet doveo je bezbroj ljudi do uspeha kako na poslovnom, tako i na privatnom planu.",
		Zanr:             model.ZanrEnum(7),
		BrojStrana:       223,
		GodinaNastanka:   1936,
		UkupnaKolicina:   8,
		TrenutnoDostupno: 8,
		ProsecnaOcena:    5.0,
		BrojOcena:        1,
		Slika:            "9788677104986.jpg",
	},
	{Isbn: "9788610030860",
		Naziv:            "Bornova Pravila",
		ImeAutora:        "Robert",
		PrezimeAutora:    "Ladlam",
		Opis:             "Pred Bornom je možda najteži izbor do sada: politička budućnost sveta ili život njemu najdražih osoba.\n\nDžejson Born suočava se sa nemogućim zadatkom. Angažovan je da glumi dvojnika visokog vladinog ministra na samitu u Kataru, kako bi ga zaštitio od pokušaja atentata. Iznenada, naoružani ljudi upadaju u prostoriju i ubijaju sve osim Borna. Njihova meta, međutim, nije ministar kog Born glumi… već Born lično.",
		Zanr:             model.ZanrEnum(4),
		BrojStrana:       368,
		GodinaNastanka:   2019,
		UkupnaKolicina:   7,
		TrenutnoDostupno: 7,
		ProsecnaOcena:    5.0,
		BrojOcena:        1,
		Slika:            "9788610030860.jpg",
	},
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
	Database.Migrator().DropTable("pretplate")

	Database.AutoMigrate(&model.Knjiga{})
	Database.AutoMigrate(&model.Pretplata{})

	for _, knjiga := range knjige {
		Database.Create(&knjiga)
	}

}
