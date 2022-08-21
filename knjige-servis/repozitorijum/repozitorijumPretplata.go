package repozitorijum

import (
	"errors"
	model "knjige-servis/model"
	util "knjige-servis/util"
	"strconv"
)

func PreuzmiSvePretplata() []model.Pretplata {
	var pretplate []model.Pretplata
	util.Database.Find(&pretplate)
	return pretplate
}

func PreuzmiPoIdPretplata(id uint) (model.Pretplata, error) {
	var pretplata model.Pretplata

	util.Database.First(&pretplata, id)

	if pretplata.ID == 0 {
		return pretplata, errors.New("Pretplata sa ID " + strconv.FormatUint(uint64(id), 10) + " ne postoji.")
	}

	return pretplata, nil
}

func PreuzmiPoKorisniku(korisnikId uint) []model.Pretplata {
	var pretplate []model.Pretplata

	util.Database.Where("korisnik_id = ?", korisnikId).Find(&pretplate)
	return pretplate
}

func PreuzmiPoKnjizi(knjigaId uint) []model.Pretplata {
	var pretplate []model.Pretplata

	util.Database.Where("knjiga_id = ?", knjigaId).Find(&pretplate)
	return pretplate
}

func KreirajPretplatu(pretplata model.Pretplata) error {
	result := util.Database.Create(&pretplata)

	return result.Error
}

func IzmeniPretplatu(pretplata model.Pretplata) error {
	result := util.Database.Save(&pretplata)

	return result.Error
}

func ObrisiPoIdPretplatu(id uint) error {
	result := util.Database.Delete(&model.Pretplata{}, id)

	return result.Error
}
