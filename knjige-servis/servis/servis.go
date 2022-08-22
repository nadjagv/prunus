package servis

import (
	"errors"
	"fmt"
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
		return errors.New("nedostaju podaci")
	} else if dto.BrojStrana <= 0 || dto.UkupnaKolicina <= 0 {
		return errors.New("broj strana i ukupna količina knjige u biblioteci moraju biti pozitivni brojevi")
	}

	dto.Slika = "default.jpg"
	dto.TrenutnoDostupno = dto.UkupnaKolicina
	dto.BrojOcena = 0
	err := repozitorijum.Kreiraj(dto.MapirajNaObjekat())

	return err
}

func Izmeni(dto model.KnjigaDTO) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(dto.Id)
	if err != nil {
		return err
	}
	if dto.GodinaNastanka == 0 {
		return errors.New("nije unesena validna godina")
	}

	if dto.BrojStrana <= 0 || dto.UkupnaKolicina <= 0 {
		return errors.New("broj strana i ukupna količina knjige u biblioteci moraju biti pozitivni brojevi")
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

func ProveriDostupnuKolicinu(id uint) (uint, error) {
	knjiga, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return 0, err
	}

	return knjiga.TrenutnoDostupno, err
}

func SmanjiDostupnuKolicinu(id uint) error {
	knjiga, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return err
	}

	if knjiga.TrenutnoDostupno < 1 {
		return errors.New("knjiga trenutno nije dostupna")
	}
	knjiga.TrenutnoDostupno -= 1
	err = repozitorijum.Izmeni(knjiga)

	return err
}

func PovecajDostupnuKolicinu(id uint) error {
	knjiga, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return err
	}

	if knjiga.TrenutnoDostupno+1 > knjiga.UkupnaKolicina {
		return errors.New("greška - svi primerci knjige su već u biblioteci")
	}

	knjiga.TrenutnoDostupno += 1

	if knjiga.TrenutnoDostupno == 1 {
		pretplate := PreuzmiPoKnjizi(id)
		for _, p := range pretplate {
			//TODO slanje mejla pretplacenim

			fmt.Println(p.KorisnikEmail)
			continue
		}
	}
	err = repozitorijum.Izmeni(knjiga)

	return err
}

func Oceni(id uint, ocena uint) error {
	knjiga, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return err
	}

	knjiga.ProsecnaOcena = (knjiga.ProsecnaOcena*float64(knjiga.BrojOcena) + float64(ocena)) / float64((knjiga.BrojOcena + 1))

	knjiga.BrojOcena += 1
	err = repozitorijum.Izmeni(knjiga)

	return err
}
