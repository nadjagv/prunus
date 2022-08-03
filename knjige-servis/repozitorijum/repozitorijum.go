package repozitorijum

import (
	"errors"
	model "knjige-servis/model"
	util "knjige-servis/util"
	"strconv"
)

func PreuzmiSve() []model.Knjiga {
	var knjige []model.Knjiga
	util.Database.Find(&knjige)
	return knjige
}

func PreuzmiPoId(id uint) (model.Knjiga, error) {
	var knjiga model.Knjiga

	util.Database.First(&knjiga, id)

	if knjiga.ID == 0 {
		return knjiga, errors.New("Knjiga sa ID " + strconv.FormatUint(uint64(id), 10) + " ne postoji.")
	}

	return knjiga, nil
}

func PreuzmiPoIsbn(isbn string) (model.Knjiga, error) {
	var knjiga model.Knjiga

	util.Database.Where("isbn = ?", isbn).First(&knjiga)

	if knjiga.ID == 0 {
		return knjiga, errors.New("Knjiga sa ISBN " + isbn + " ne postoji.")
	}

	return knjiga, nil
}

func Kreiraj(knjiga model.Knjiga) error {
	result := util.Database.Create(&knjiga)

	return result.Error
}

func Izmeni(knjiga model.Knjiga) error {
	result := util.Database.Save(&knjiga)

	return result.Error
}

func Obrisi(knjiga model.Knjiga) error {
	result := util.Database.Delete(&knjiga)

	return result.Error
}

func ObrisiPoId(id uint) error {
	result := util.Database.Delete(&model.Knjiga{}, id)

	return result.Error
}
