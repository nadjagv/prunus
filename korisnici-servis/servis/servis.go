package servis

import (
	"errors"
	"fmt"
	model "korisnici-servis/model"
	repozitorijum "korisnici-servis/repozitorijum"
	util "korisnici-servis/util"

	"golang.org/x/crypto/bcrypt"
)

func ProveriKredencijale(kredencijali util.Kredencijali) (model.Korisnik, error) {
	korisnikUBazi, err := repozitorijum.PreuzmiPoEmail(kredencijali.Email)
	if err != nil {
		return model.Korisnik{}, errors.New("Nepostojeća email adresa.")
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
