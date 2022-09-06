package servis

import (
	"encoding/json"
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

func PreuzmiPoslednjih5KorisnikIzn(korisnikId uint) []model.Iznajmljivanje {
	sve := repozitorijum.PreuzmiSveKorisnikIzn(korisnikId)
	if len(sve) > 5 {
		return sve[len(sve)-5:]
	}
	return sve
}

func PreuzmiIzmedjuDatumaIzn(d1 time.Time, d2 time.Time) []model.Iznajmljivanje {
	return repozitorijum.PreuzmiIzmedjuDatumaIzn(d1, d2)
}

func KreirajIzn(dto model.NovoIznajmljivanjeDTO) error {
	fmt.Println(dto)
	if dto.KnjigaId == 0 || dto.Email == " " {

		return errors.New("nedostaju podaci")
	}
	//provera da li korisnik postoji
	response, err := http.Get(korisniciServisUrl + "email/" + dto.Email)
	if err != nil {
		return errors.New("korisnik ne postoji")
	}

	var korisnik model.KorisnikDTO
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&korisnik)

	//upit da se smanji dostupna kolicina
	postojiRezervacija := false
	rezervacijeKorisnika := repozitorijum.PreuzmiAktivneKorisnikRez(korisnik.Id)
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

	var novi model.Iznajmljivanje
	novi.KnjigaId = dto.KnjigaId
	novi.KorisnikId = korisnik.Id
	novi.DatumVremeIznajmljivanja = time.Now()
	novi.RokVracanja = time.Now().AddDate(0, 0, 14)
	novi.ZakasneloVracanje = false
	novi.Produzeno = false
	novi.Aktivno = true
	err = repozitorijum.KreirajIzn(novi)

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
