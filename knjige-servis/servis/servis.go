package servis

import (
	"errors"
	model "knjige-servis/model"
	repozitorijum "knjige-servis/repozitorijum"
	"strconv"
)

func PreuzmiSve() []model.Knjiga {
	return repozitorijum.PreuzmiSve()
}

func PreuzmiPoId(id uint) (model.Knjiga, error) {
	return repozitorijum.PreuzmiPoId(id)
}

func PreuzmiPoIsbn(isbn string) (model.Knjiga, error) {
	return repozitorijum.PreuzmiPoIsbn(isbn)
}

func Kreiraj(dto model.KnjigaDTO) error {
	if dto.Isbn == "" || dto.Naziv == "" || dto.ImeAutora == "" || dto.PrezimeAutora == "" || dto.Opis == "" || dto.BrojStrana == 0 || dto.GodinaNastanka == 0 || dto.UkupnaKolicina == 0 {
		return errors.New("Nedostaju podaci.")
	} else if dto.BrojStrana <= 0 || dto.UkupnaKolicina <= 0 {
		return errors.New("Broj strana i ukupna količina knjige u biblioteci moraju biti pozitivni brojevi.")
	}

	dto.Slika = "default.jpg"
	dto.TrenutnoDostupno = dto.UkupnaKolicina
	err := repozitorijum.Kreiraj(dto.MapirajNaObjekat())

	return err
}

func Izmeni(dto model.KnjigaDTO) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(dto.Id)
	if err != nil {
		return err
	}
	if dto.GodinaNastanka == 0 {
		return errors.New("Nije unesena validna godina.")
	}

	if dto.BrojStrana <= 0 || dto.UkupnaKolicina <= 0 {
		return errors.New("Broj strana i ukupna količina knjige u biblioteci moraju biti pozitivni brojevi.")
	}

	zaIzmenu.Isbn = dto.Isbn
	zaIzmenu.Naziv = dto.Naziv
	zaIzmenu.ImeAutora = dto.ImeAutora
	zaIzmenu.PrezimeAutora = dto.PrezimeAutora
	zaIzmenu.Opis = dto.Opis
	zaIzmenu.BrojStrana = dto.BrojStrana
	zaIzmenu.GodinaNastanka = dto.GodinaNastanka
	zaIzmenu.Zanr = dto.Zanr

	staraUkupnaKolicina := zaIzmenu.UkupnaKolicina
	brojZauzetih := staraUkupnaKolicina - zaIzmenu.TrenutnoDostupno
	if dto.UkupnaKolicina < brojZauzetih {
		return errors.New("Nije moguće postaviti ukupnu količinu na " + strconv.FormatUint(uint64(dto.UkupnaKolicina), 10) + ", broj zauzetih knjiga je: " + strconv.FormatUint(uint64(brojZauzetih), 10))
	}
	zaIzmenu.UkupnaKolicina = dto.UkupnaKolicina
	zaIzmenu.TrenutnoDostupno = zaIzmenu.TrenutnoDostupno - (staraUkupnaKolicina - dto.UkupnaKolicina)

	err = repozitorijum.Izmeni(zaIzmenu)
	return err
}

func ObrisiPoId(id uint) error {
	return repozitorijum.ObrisiPoId(id)
}
