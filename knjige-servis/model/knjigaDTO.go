package model

type KnjigaDTO struct {
	Id               uint
	Isbn             string   `gorm:"not null;"`
	Naziv            string   `gorm:"not null;"`
	ImeAutora        string   `gorm:"not null;"`
	PrezimeAutora    string   `gorm:"not null;"`
	Opis             string   `gorm:"not null;"`
	Zanr             ZanrEnum `gorm:"not null;default:12"`
	BrojStrana       uint     `gorm:"not null;"`
	GodinaNastanka   uint     `gorm:"not null;"`
	UkupnaKolicina   uint     `gorm:"not null;"`
	TrenutnoDostupno uint     `gorm:"not null;"`
	ProsecnaOcena    float64  `gorm:"not null;"`
	Slika            string
}

func (knjiga *Knjiga) MapirajNaDTO() KnjigaDTO {
	return KnjigaDTO{
		Id:               knjiga.ID,
		Isbn:             knjiga.Isbn,
		Naziv:            knjiga.Naziv,
		ImeAutora:        knjiga.ImeAutora,
		PrezimeAutora:    knjiga.PrezimeAutora,
		Opis:             knjiga.Opis,
		Zanr:             knjiga.Zanr,
		BrojStrana:       knjiga.BrojStrana,
		GodinaNastanka:   knjiga.GodinaNastanka,
		UkupnaKolicina:   knjiga.UkupnaKolicina,
		TrenutnoDostupno: knjiga.TrenutnoDostupno,
		ProsecnaOcena:    knjiga.ProsecnaOcena,
		Slika:            knjiga.Slika,
	}
}

func (knjiga *KnjigaDTO) MapirajNaObjekat() Knjiga {
	return Knjiga{
		Isbn:             knjiga.Isbn,
		Naziv:            knjiga.Naziv,
		ImeAutora:        knjiga.ImeAutora,
		PrezimeAutora:    knjiga.PrezimeAutora,
		Opis:             knjiga.Opis,
		Zanr:             knjiga.Zanr,
		BrojStrana:       knjiga.BrojStrana,
		GodinaNastanka:   knjiga.GodinaNastanka,
		UkupnaKolicina:   knjiga.UkupnaKolicina,
		TrenutnoDostupno: knjiga.TrenutnoDostupno,
		ProsecnaOcena:    knjiga.ProsecnaOcena,
		Slika:            knjiga.Slika,
	}
}
