package servis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	model "korisnici-servis/model"
	repozitorijum "korisnici-servis/repozitorijum"
	util "korisnici-servis/util"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ProveriKredencijale(kredencijali util.Kredencijali) (model.Korisnik, error) {
	korisnikUBazi, err := repozitorijum.PreuzmiPoEmail(kredencijali.Email)
	if err != nil {
		return model.Korisnik{}, errors.New("Nepostojeća email adresa.")
	}

	if korisnikUBazi.Blokiran {
		return model.Korisnik{}, errors.New("Korisnik blokiran.")
	}
	ocekivanaLozinka := korisnikUBazi.Lozinka

	err = bcrypt.CompareHashAndPassword([]byte(ocekivanaLozinka), []byte(kredencijali.Lozinka))
	if err != nil {
		fmt.Println("ovde")
		return model.Korisnik{}, errors.New("Pogrešna lozinka.")
	}

	// if string(hashLozinka) != ocekivanaLozinka {
	// 	return errors.New("Pogrešna lozinka.")
	// }

	return korisnikUBazi, nil
}

func PreuzmiSve() []model.Korisnik {
	return repozitorijum.PreuzmiSve()
}

func PreuzmiPoId(id uint) (model.Korisnik, error) {
	return repozitorijum.PreuzmiPoId(id)
}

func PreuzmiPoEmail(email string) (model.Korisnik, error) {
	return repozitorijum.PreuzmiPoEmail(email)
}

func Kreiraj(dto model.KorisnikDTO) error {
	if dto.Email == "" || dto.Ime == "" || dto.Prezime == "" || dto.Lozinka == "" {
		return errors.New("Nedostaju podaci.")
	}

	if !util.ValidanEmail(dto.Email) {
		return errors.New("Mejl adresa nije validna.")
	}

	if !util.ValidnaLozinka(dto.Lozinka) {
		fmt.Println(dto.Lozinka)
		return errors.New("Lozinka nije validna.")
	}

	_, err := repozitorijum.PreuzmiPoEmail(dto.Email)
	if err == nil {
		return errors.New("Nalog sa istom mejl adresom već postoji.")

	}
	hesirano, _ := util.Hesiraj(dto.Lozinka)

	dto.Lozinka = string(hesirano)

	dto.Sumnjiv = false
	dto.Blokiran = false
	dto.IstekClanarine = time.Now()
	err = repozitorijum.Kreiraj(dto.MapirajNaObjekat())

	return err
}

func Izmeni(dto model.KorisnikDTO) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(dto.Id)
	if err != nil {
		return err
	}
	if !util.ValidanEmail(dto.Email) {
		return errors.New("Mejl adresa nije validna.")
	}
	zaIzmenu.Ime = dto.Ime
	zaIzmenu.Prezime = dto.Prezime

	err = repozitorijum.Izmeni(zaIzmenu)
	return err
}

func IzmeniLozinku(dto model.IzmenaLozinkeDTO) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(dto.KorisnikId)
	if err != nil {
		return err
	}

	ocekivanaLozinka := zaIzmenu.Lozinka

	err = bcrypt.CompareHashAndPassword([]byte(ocekivanaLozinka), []byte(dto.Stara))
	if err != nil {
		return errors.New("Pogrešna lozinka.")
	}

	if !util.ValidnaLozinka(dto.Nova) {
		return errors.New("Lozinka nije validna.")
	}

	hesirano, _ := util.Hesiraj(dto.Nova)

	zaIzmenu.Lozinka = string(hesirano)

	err = repozitorijum.Izmeni(zaIzmenu)
	return err
}

func ObrisiPoId(id uint) error {
	return repozitorijum.ObrisiPoId(id)
}

func OznaciSumnjiv(id uint) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return err
	}
	zaIzmenu.Sumnjiv = true

	poruka := "Opomena: primetili smo da ste zakasnili sa vraćanjem knjige 3 puta u proteklih 6 meseci. Molimo Vas da vraćate knjige na vreme."
	mejl := model.Mejl{
		Poruka:     poruka,
		MejlAdresa: zaIzmenu.Email,
	}

	jsonMejl, err := json.Marshal(mejl)
	if err != nil {
		return err
	}
	_, err = http.Post("http://localhost:8084/opomena", "application/json", bytes.NewReader([]byte(jsonMejl)))
	if err != nil {
		return err
	}

	err = repozitorijum.Izmeni(zaIzmenu)

	return err
}

func ProduziClanarinu(id uint) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return err
	}

	if zaIzmenu.IstekClanarine.Before(time.Now()) {
		zaIzmenu.IstekClanarine = time.Now().AddDate(1, 0, 0)
	} else {
		zaIzmenu.IstekClanarine = zaIzmenu.IstekClanarine.AddDate(1, 0, 0)
	}

	err = repozitorijum.Izmeni(zaIzmenu)
	return err
}

func Blokiraj(id uint, obrazlozenje string) error {
	zaIzmenu, err := repozitorijum.PreuzmiPoId(id)
	if err != nil {
		return err
	}

	zaIzmenu.Blokiran = true

	err = repozitorijum.Izmeni(zaIzmenu)

	poruka := "Nalog blokiran. Obrazloženje: \n" + obrazlozenje
	mejl := model.Mejl{
		Poruka:     poruka,
		MejlAdresa: zaIzmenu.Email,
	}

	jsonMejl, err := json.Marshal(mejl)
	if err != nil {
		return err
	}
	_, err = http.Post("http://localhost:8084/blokiranje", "application/json", bytes.NewReader([]byte(jsonMejl)))
	if err != nil {
		return err
	}
	return err
}
