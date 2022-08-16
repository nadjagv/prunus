package repozitorijum

import (
	"errors"
	model "korisnici-servis/model"
	util "korisnici-servis/util"
	"strconv"
)

func PreuzmiSve() []model.Korisnik {
	var korisnici []model.Korisnik
	util.Database.Find(&korisnici)
	return korisnici
}

func PreuzmiPoId(id uint) (model.Korisnik, error) {
	var korisnik model.Korisnik

	util.Database.First(&korisnik, id)

	if korisnik.ID == 0 {
		return korisnik, errors.New("Korisnik sa ID " + strconv.FormatUint(uint64(id), 10) + " ne postoji.")
	}

	return korisnik, nil
}

func PreuzmiPoEmail(email string) (model.Korisnik, error) {
	var korisnik model.Korisnik

	util.Database.Where("email = ?", email).First(&korisnik)

	if korisnik.ID == 0 {
		return korisnik, errors.New("Korisnik sa email adresom " + email + " ne postoji.")
	}

	return korisnik, nil
}

func Kreiraj(korisnik model.Korisnik) error {
	result := util.Database.Create(&korisnik)

	return result.Error
}

func Izmeni(korisnik model.Korisnik) error {
	result := util.Database.Save(&korisnik)

	return result.Error
}

func Obrisi(korisnik model.Korisnik) error {
	result := util.Database.Delete(&korisnik)

	return result.Error
}

func ObrisiPoId(id uint) error {
	result := util.Database.Delete(&model.Korisnik{}, id)

	return result.Error
}
