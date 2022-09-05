package repozitorijum

import (
	"errors"
	model "rezervacija-iznajmljivanje-servis/model"
	util "rezervacija-iznajmljivanje-servis/util"
	"strconv"
	"time"
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

func PreuzmiAktivnaIzn() []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje

	util.Database.Where("aktivno = ?", true).Find(&iznajmljivanja)

	return iznajmljivanja
}

func PreuzmiSveKorisnikIzn(id uint) []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje

	util.Database.Where("korisnik_id = ?", id).Find(&iznajmljivanja)

	return iznajmljivanja
}

func PreuzmiAktivnaKorisnikIzn(id uint) []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje

	util.Database.Where("korisnik_id = ? AND aktivno = ?", id, true).Find(&iznajmljivanja)

	return iznajmljivanja
}

func PreuzmiAktivnuKorisnikKnjigaIzn(id uint, knjigaId uint) model.Iznajmljivanje {
	var iznajmljivanje model.Iznajmljivanje

	util.Database.Where("korisnik_id = ? AND aktivno = ? AND knjiga_id=?", id, true, knjigaId).First(&iznajmljivanje)

	return iznajmljivanje
}

func PreuzmiIzmedjuDatumaZaKorisnikaIzn(korisnikId uint, d1 time.Time, d2 time.Time) []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje

	util.Database.Where("korisnik_id = ? AND datum_vreme_iznajmljivanja BETWEEN ? AND ?", korisnikId, d1, d2).Find(&iznajmljivanja)

	return iznajmljivanja
}

func PreuzmiIzmedjuDatumaIzn(d1 time.Time, d2 time.Time) []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje

	util.Database.Where("datum_vreme_iznajmljivanja BETWEEN ? AND ?", d1, d2).Find(&iznajmljivanja)

	return iznajmljivanja
}

func PreuzmiZakasneleIzmedjuDatumaZaKorisnikaIzn(korisnikId uint, d1 time.Time, d2 time.Time) []model.Iznajmljivanje {
	var iznajmljivanja []model.Iznajmljivanje

	util.Database.Where("korisnik_id = ? AND zakasnelo_vracanje=? AND datum_vreme_iznajmljivanja BETWEEN ? AND ?", korisnikId, true, d1, d2).Find(&iznajmljivanja)

	return iznajmljivanja
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
