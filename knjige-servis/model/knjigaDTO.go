package model

type KnjigaDTO struct {
	Id               uint
	Isbn             string
	Naziv            string
	ImeAutora        string
	PrezimeAutora    string
	Opis             string
	Zanr             ZanrEnum
	BrojStrana       uint
	GodinaNastanka   uint
	UkupnaKolicina   uint
	TrenutnoDostupno uint
	ProsecnaOcena    float64
	BrojOcena        uint
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
		BrojOcena:        knjiga.BrojOcena,
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
		BrojOcena:        knjiga.BrojOcena,
		Slika:            knjiga.Slika,
	}
}
