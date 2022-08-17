package servis

import (
	"errors"
	"fmt"
	"net/http"
	model "rezervacija-iznajmljivanje-servis/model"
	repozitorijum "rezervacija-iznajmljivanje-servis/repozitorijum"
	"strconv"
	"time"
)

func PreuzmiSveRez() []model.Rezervacija {
	return repozitorijum.PreuzmiSveRez()
}

func PreuzmiPoIdRez(id uint) (model.Rezervacija, error) {
	return repozitorijum.PreuzmiPoIdRez(id)
}

func KreirajRez(dto model.RezervacijaDTO) error {
	if dto.KnjigaId == 0 || dto.KorisnikId == 0 {
		return errors.New("Nedostaju podaci.")
	}
	//provera da li korisnik postoji
	_, err := http.Get(korisniciServisUrl + strconv.FormatUint(uint64(dto.KorisnikId), 10))
	if err != nil {
		fmt.Println(err)
		return errors.New("Korisnik ne postoji")
	}

	//upit da se smanji dostupna kolicina
	request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"smanji-kolicinu/"+strconv.FormatUint(uint64(dto.KnjigaId), 10), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return errors.New("Nije moguće izvršiti - greška pri smanjivanju količine knjige.")
	}

	dto.DatumVremeIsteka = time.Now().AddDate(0, 0, 7) //7 dana traje rezervacija
	err = repozitorijum.KreirajRez(dto.MapirajNaObjekat())

	return err
}

func ObrisiPoIdRez(id uint) error {
	return repozitorijum.ObrisiPoIdRez(id)
}
