package servis

import (
	"errors"
	"fmt"
	"net/http"
	model "rezervacija-iznajmljivanje-servis/model"
	repozitorijum "rezervacija-iznajmljivanje-servis/repozitorijum"
	"strconv"
	"time"

	"github.com/robfig/cron"
)

func PreuzmiSveRez() []model.Rezervacija {
	return repozitorijum.PreuzmiSveRez()
}

func PreuzmiPoIdRez(id uint) (model.Rezervacija, error) {
	return repozitorijum.PreuzmiPoIdRez(id)
}
func PreuzmiAktivnuKorisnikKnjigaRez(id uint, knjigaId uint) model.Rezervacija {
	return repozitorijum.PreuzmiAktivnuKorisnikKnjigaRez(id, knjigaId)
}

func PreuzmiAktivneKorisnikRez(id uint) []model.Rezervacija {
	return repozitorijum.PreuzmiAktivneKorisnikRez(id)
}

func KreirajRez(dto model.RezervacijaDTO) error {
	if dto.KnjigaId == 0 || dto.KorisnikId == 0 {
		return errors.New("nedostaju podaci")
	}
	//provera da li korisnik postoji
	_, err := http.Get(korisniciServisUrl + strconv.FormatUint(uint64(dto.KorisnikId), 10))
	if err != nil {
		fmt.Println(err)
		return errors.New("korisnik ne postoji")
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
		return errors.New("nije moguće izvršiti - greška pri smanjivanju količine knjige")
	}

	dto.DatumVremeIsteka = time.Now().AddDate(0, 0, 7) //7 dana traje rezervacija
	dto.Aktivno = true
	err = repozitorijum.KreirajRez(dto.MapirajNaObjekat())

	return err
}

func OtkaziRezervaciju(id uint) error {
	rezervacija, err := repozitorijum.PreuzmiPoIdRez(id)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"povecaj-kolicinu/"+strconv.FormatUint(uint64(id), 10), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return errors.New("nije moguće izvršiti - greška pri povećavanju količine knjige")
	}

	rezervacija.Aktivno = false
	err = repozitorijum.IzmeniRez(rezervacija)
	return err
}

func ProveravajIstekRezervacija() {
	customLocation, _ := time.LoadLocation("Europe/Belgrade")
	cronHandler := cron.NewWithLocation(customLocation)

	cronHandler.AddFunc("@midnight", func() {
		aktivneRezervacije := repozitorijum.PreuzmiAktivneRez()
		for _, r := range aktivneRezervacije {
			if r.DatumVremeIsteka.Before(time.Now()) {
				OtkaziRezervaciju(r.ID)
			}
		}
	})
}
func ObrisiPoIdRez(id uint) error {
	return repozitorijum.ObrisiPoIdRez(id)
}
