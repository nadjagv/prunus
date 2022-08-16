package model

import "time"

type KorisnikDTO struct {
	Id             uint
	Email          string
	Lozinka        string
	Ime            string
	Prezime        string
	Tip            TipKorisnika
	IstekClanarine time.Time
	Sumnjiv        bool
	Blokiran       bool
}

func (korisnik *Korisnik) MapirajNaDTO() KorisnikDTO {
	return KorisnikDTO{
		Id:             korisnik.ID,
		Email:          korisnik.Email,
		Lozinka:        korisnik.Lozinka,
		Ime:            korisnik.Ime,
		Prezime:        korisnik.Prezime,
		Tip:            korisnik.Tip,
		IstekClanarine: korisnik.IstekClanarine,
		Sumnjiv:        korisnik.Sumnjiv,
		Blokiran:       korisnik.Blokiran,
	}
}

func (korisnik *KorisnikDTO) MapirajNaObjekat() Korisnik {
	return Korisnik{
		Email:          korisnik.Email,
		Lozinka:        korisnik.Lozinka,
		Ime:            korisnik.Ime,
		Prezime:        korisnik.Prezime,
		Tip:            korisnik.Tip,
		IstekClanarine: korisnik.IstekClanarine,
		Sumnjiv:        korisnik.Sumnjiv,
		Blokiran:       korisnik.Blokiran,
	}
}
