package repozitorijum

import (
	"errors"
	model "rezervacija-iznajmljivanje-servis/model"
	util "rezervacija-iznajmljivanje-servis/util"
	"strconv"
)

func PreuzmiSveIzn() []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje
	util.Database.Find(&iznajmljivanja)
	return iznajmljivanja
}

func PreuzmiPoIdIzn(id uint) (model.Iznajmljivanje, error) {
	var iznajmljivanje model.Iznajmljivanje

	util.Database.First(&iznajmljivanje, id)

	if iznajmljivanje.ID == 0 {
		return iznajmljivanje, errors.New("Iznajmljivanje sa ID " + strconv.FormatUint(uint64(id), 10) + " ne postoji.")
	}

	return iznajmljivanje, nil
}

func KreirajIzn(iznajmljivanje model.Iznajmljivanje) error {
	result := util.Database.Create(&iznajmljivanje)

	return result.Error
}

func IzmeniIzn(iznajmljivanje model.Iznajmljivanje) error {
	result := util.Database.Save(&iznajmljivanje)

	return result.Error
}

func ObrisiIzn(iznajmljivanje model.Iznajmljivanje) error {
	result := util.Database.Delete(&iznajmljivanje)

	return result.Error
}

func ObrisiPoIdIzn(id uint) error {
	result := util.Database.Delete(&model.Iznajmljivanje{}, id)

	return result.Error
}
