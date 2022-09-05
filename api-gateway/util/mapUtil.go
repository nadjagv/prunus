package util

import (
	"net/http"
	"strconv"

	"api-gateway/dto"
)

func MapirajNaRecenzijeNazivEmailDTO(rec dto.RecenzijaDTO) (dto.RecenzijaNazivEmailDTO, error) {
	var korisniciServisUrl = "http://localhost:8082/"
	var knjigeServisUrl = "http://localhost:8081/"
	var knjiga dto.KnjigaDTO
	var korisnik dto.KorisnikDTO
	response, err := http.Get(knjigeServisUrl + strconv.FormatUint(uint64(rec.Knjiga_id), 10))
	if err != nil {
		return dto.RecenzijaNazivEmailDTO{}, err
	}

	err = GetJson(response, &knjiga)
	if err != nil {
		return dto.RecenzijaNazivEmailDTO{}, err
	}

	response, err = http.Get(korisniciServisUrl + strconv.FormatUint(uint64(rec.Korisnik_id), 10))
	if err != nil {
		return dto.RecenzijaNazivEmailDTO{}, err
	}

	err = GetJson(response, &korisnik)
	if err != nil {
		return dto.RecenzijaNazivEmailDTO{}, err
	}

	rezDto := dto.RecenzijaNazivEmailDTO{
		Id:            rec.Id,
		KorisnikId:    rec.Korisnik_id,
		KnjigaId:      rec.Knjiga_id,
		KnjigaNaziv:   knjiga.Naziv,
		Komentar:      rec.Komentar,
		Ocena:         rec.Ocena,
		Status:        rec.Status,
		Obrisano:      rec.Obrisano,
		KorisnikEmail: korisnik.Email,
	}
	return rezDto, nil
}
