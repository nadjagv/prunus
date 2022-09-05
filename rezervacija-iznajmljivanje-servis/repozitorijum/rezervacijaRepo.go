package repozitorijum

import (
	"errors"
	model "rezervacija-iznajmljivanje-servis/model"
	util "rezervacija-iznajmljivanje-servis/util"
	"strconv"
)

func PreuzmiSveRez() []model.Rezervacija {
	var rezervacije []model.Rezervacija
	util.Database.Find(&rezervacije)
	return rezervacije
}

func PreuzmiPoIdRez(id uint) (model.Rezervacija, error) {
	var rezervacija model.Rezervacija

	util.Database.First(&rezervacija, id)

	if rezervacija.ID == 0 {
		return rezervacija, errors.New("Rezervacija sa ID " + strconv.FormatUint(uint64(id), 10) + " ne postoji.")
	}

	return rezervacija, nil
}

func PreuzmiAktivneRez() []model.Rezervacija {
	var rezervacije []model.Rezervacija

	util.Database.Where("aktivno = ?", true).Find(&rezervacije)

	return rezervacije
}

func PreuzmiSveKorisnikRez(id uint) []model.Rezervacija {
	var rezervacije []model.Rezervacija

	util.Database.Where("korisnik_id = ?", id).Find(&rezervacije)

	return rezervacije
}

func PreuzmiAktivneKorisnikRez(id uint) []model.Rezervacija {
	var rezervacije []model.Rezervacija

	util.Database.Where("korisnik_id = ? AND aktivno = ?", id, true).Find(&rezervacije)

	return rezervacije
}

func PreuzmiAktivnuKorisnikKnjigaRez(id uint, knjigaId uint) model.Rezervacija {
	var rezervacija model.Rezervacija

	util.Database.Where("korisnik_id = ? AND aktivno = ? AND knjiga_id=?", id, true, knjigaId).First(&rezervacija)

	return rezervacija
}

func KreirajRez(rezervacija model.Rezervacija) error {
	result := util.Database.Create(&rezervacija)

	return result.Error
}

func IzmeniRez(rezervacija model.Rezervacija) error {
	result := util.Database.Save(&rezervacija)

	return result.Error
}

func ObrisiRez(rezervacija model.Rezervacija) error {
	result := util.Database.Delete(&rezervacija)

	return result.Error
}

func ObrisiPoIdRez(id uint) error {
	result := util.Database.Delete(&model.Rezervacija{}, id)

	return result.Error
}
