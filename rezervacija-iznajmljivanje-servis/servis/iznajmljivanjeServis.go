package servis

import (
	"errors"
	"net/http"
	model "rezervacija-iznajmljivanje-servis/model"
	repozitorijum "rezervacija-iznajmljivanje-servis/repozitorijum"
	"strconv"
	"time"
)

var knjigeServisUrl = "http://localhost:8081/"
var korisniciServisUrl = "http://localhost:8082/"

func PreuzmiSveIzn() []model.Iznajmljivanje {
	return repozitorijum.PreuzmiSveIzn()
}

func PreuzmiPoIdIzn(id uint) (model.Iznajmljivanje, error) {
	return repozitorijum.PreuzmiPoIdIzn(id)
}

func KreirajIzn(dto model.IznajmljivanjeDTO) error {
	if dto.KnjigaId == 0 || dto.KorisnikId == 0 {
		return errors.New("Nedostaju podaci")
	}
	//provera da li korisnik postoji
	_, err := http.Get(korisniciServisUrl + strconv.FormatUint(uint64(dto.KorisnikId), 10))
	if err != nil {
		return errors.New("Korisnik ne postoji")
	}

	//upit da se smanji dostupna kolicina
	request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"smanji-kolicinu/"+strconv.FormatUint(uint64(dto.KnjigaId), 10), nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return errors.New("Nije moguće izvršiti - greška pri smanjivanju količine knjige.")
	}

	dto.DatumVremeIznajmljivanja = time.Now()
	dto.RokVracanja = time.Now().AddDate(0, 0, 14)
	dto.ZakasneloVracanje = false
	dto.Produzeno = false
	err = repozitorijum.KreirajIzn(dto.MapirajNaObjekat())

	return err
}

func ObrisiPoIdIzn(id uint) error {
	return repozitorijum.ObrisiPoIdIzn(id)
}
