package servis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	model "knjige-servis/model"
	repozitorijum "knjige-servis/repozitorijum"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	if dto.Slika == "" {
		dto.Slika = "default.jpg"
	} else {
		base64Slika := dto.Slika
		indeks := strings.Index(string(base64Slika), ",")
		mimeType := strings.TrimSuffix(base64Slika[5:indeks], ";base64")

		putanja := ""
		switch mimeType {
		case "image/jpeg":
			putanja = dto.Isbn + ".jpg"
		case "image/png":
			putanja = dto.Isbn + ".png"
		}
		sep := string(os.PathSeparator)
		model.KonvertujIzBase64USliku(base64Slika, "slike"+sep+putanja)
		dto.Slika = putanja
	}
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

		poruka := "Obaveštavamo Vas da je knjiga " + knjiga.Naziv + " ponovo dostupna."
		mejl := model.Mejl{
			Poruka:     poruka,
			MejlAdresa: "",
		}
		for _, p := range pretplate {
			//slanje mejla pretplacenim
			mejl.MejlAdresa = p.KorisnikEmail

			jsonMejl, err := json.Marshal(mejl)
			if err != nil {
				return err
			}
			_, err = http.Post("http://localhost:8084/pretplata", "application/json", bytes.NewReader([]byte(jsonMejl)))
			if err != nil {
				return err
			}

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
