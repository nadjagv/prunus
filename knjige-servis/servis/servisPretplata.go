package servis

import (
	"errors"
	model "knjige-servis/model"
	repozitorijum "knjige-servis/repozitorijum"
	"net/http"
	"strconv"
)

func PreuzmiSvePretplata() []model.Pretplata {
	return repozitorijum.PreuzmiSvePretplata()
}

func PreuzmiPoIdPretplata(id uint) (model.Pretplata, error) {
	return repozitorijum.PreuzmiPoIdPretplata(id)
}

func PreuzmiPoKorisniku(korisnikId uint) []model.Pretplata {
	return repozitorijum.PreuzmiPoKorisniku(korisnikId)
}

func PreuzmiPoKnjizi(knjigaId uint) []model.Pretplata {
	return repozitorijum.PreuzmiPoKnjizi(knjigaId)
}

func PreuzmiPoKnjiziKorisniku(knjigaId uint, korisnikId uint) model.Pretplata {
	return repozitorijum.PreuzmiPoKnjiziKorisniku(knjigaId, korisnikId)
}

func KreirajPretplatu(pretplata model.Pretplata) error {
	if pretplata.KnjigaId == 0 || pretplata.KorisnikId == 0 || pretplata.KorisnikEmail == "" {
		return errors.New("nedostaju podaci")
	}

	_, err := repozitorijum.PreuzmiPoId(pretplata.KnjigaId)
	if err != nil {
		return err
	}

	_, err = http.Get("http://localhost:8082/" + strconv.FormatUint(uint64(pretplata.ID), 10))
	if err != nil {
		return errors.New("korisnik ne postoji")
	}

	pretplateKorisnika := PreuzmiPoKorisniku(pretplata.KorisnikId)
	for _, p := range pretplateKorisnika {
		if p.KnjigaId == pretplata.KnjigaId {
			return errors.New("korisnik je vec pretplacen na ovu knjigu")
		}
	}

	err = repozitorijum.KreirajPretplatu(pretplata)

	return err

}

func ObrisiPoIdPretplatu(id uint) error {
	return repozitorijum.ObrisiPoIdPretplatu(id)
}
