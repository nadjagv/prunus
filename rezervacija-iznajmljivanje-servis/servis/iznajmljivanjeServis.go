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

var knjigeServisUrl = "http://localhost:8081/"
var korisniciServisUrl = "http://localhost:8082/"

func PreuzmiSveIzn() []model.Iznajmljivanje {
	return repozitorijum.PreuzmiSveIzn()
}

func PreuzmiPoIdIzn(id uint) (model.Iznajmljivanje, error) {
	return repozitorijum.PreuzmiPoIdIzn(id)
}

func PreuzmiAktivnuKorisnikKnjigaIzn(id uint, knjigaId uint) model.Iznajmljivanje {
	return repozitorijum.PreuzmiAktivnuKorisnikKnjigaIzn(id, knjigaId)
}

func PreuzmiPoKorisnikuAktivnaIzn(korisnikId uint) []model.Iznajmljivanje {
	return repozitorijum.PreuzmiAktivnaKorisnikIzn(korisnikId)
}

func KreirajIzn(dto model.IznajmljivanjeDTO) error {
	if dto.KnjigaId == 0 || dto.KorisnikId == 0 {
		return errors.New("nedostaju podaci")
	}
	//provera da li korisnik postoji
	_, err := http.Get(korisniciServisUrl + strconv.FormatUint(uint64(dto.KorisnikId), 10))
	if err != nil {
		return errors.New("korisnik ne postoji")
	}

	//upit da se smanji dostupna kolicina
	postojiRezervacija := false
	rezervacijeKorisnika := repozitorijum.PreuzmiAktivneKorisnikRez(dto.KorisnikId)
	for _, r := range rezervacijeKorisnika {
		if r.KnjigaId == dto.KnjigaId {
			postojiRezervacija = true
			r.Aktivno = false
			repozitorijum.IzmeniRez(r)
			break
		}
	}

	if !postojiRezervacija {
		request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"smanji-kolicinu/"+strconv.FormatUint(uint64(dto.KnjigaId), 10), nil)
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		client := &http.Client{}
		_, err = client.Do(request)
		if err != nil {
			return errors.New("nije moguće izvršiti - greška pri smanjivanju količine knjige")
		}
	}

	dto.DatumVremeIznajmljivanja = time.Now()
	dto.RokVracanja = time.Now().AddDate(0, 0, 14)
	dto.ZakasneloVracanje = false
	dto.Produzeno = false
	dto.Aktivno = true
	err = repozitorijum.KreirajIzn(dto.MapirajNaObjekat())

	return err
}

func VratiKnjigu(dto model.IznajmljivanjeDTO) error {
	iznajmljivanje, err := repozitorijum.PreuzmiPoIdIzn(dto.Id)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, knjigeServisUrl+"povecaj-kolicinu/"+strconv.FormatUint(uint64(dto.KnjigaId), 10), nil)
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

	iznajmljivanje.Aktivno = false
	iznajmljivanje.DatumVremeVracanja = time.Now()
	if iznajmljivanje.RokVracanja.Before(iznajmljivanje.DatumVremeVracanja) {
		iznajmljivanje.ZakasneloVracanje = true

		zakasnelaVracanja := repozitorijum.PreuzmiZakasneleIzmedjuDatumaZaKorisnikaIzn(iznajmljivanje.KorisnikId, time.Now().AddDate(0, -6, 0), time.Now()) //unazad 6 meseci
		if len(zakasnelaVracanja) >= 3 {
			//upit za oznacavanje da je sumnjiv
			request, err := http.NewRequest(http.MethodPut, korisniciServisUrl+"sumnjiv/"+strconv.FormatUint(uint64(iznajmljivanje.KorisnikId), 10), nil)
			if err != nil {
				fmt.Println(err)
				return err
			}
			request.Header.Set("Content-Type", "application/json; charset=utf-8")
			client := &http.Client{}
			_, err = client.Do(request)
			if err != nil {
				fmt.Println(err)
				return errors.New("nije moguće izvršiti - greška pri oznacavanju da je korisnik sumnjiv")
			}
		}
	} else {
		iznajmljivanje.ZakasneloVracanje = false
	}
	err = repozitorijum.IzmeniIzn(iznajmljivanje)
	return err
}

func ProduziIzn(id uint) error {
	iznajmljivanje, err := repozitorijum.PreuzmiPoIdIzn(id)
	if err != nil {
		return err
	}

	if iznajmljivanje.Produzeno {
		return errors.New("vec je produzeno")
	}

	if !iznajmljivanje.Aktivno {
		return errors.New("nije aktivno")
	}

	iznajmljivanje.RokVracanja = iznajmljivanje.RokVracanja.AddDate(0, 0, 7) //produzenje za 7 dana
	iznajmljivanje.Produzeno = true

	repozitorijum.IzmeniIzn(iznajmljivanje)

	return nil
}

func ObrisiPoIdIzn(id uint) error {
	return repozitorijum.ObrisiPoIdIzn(id)
}
