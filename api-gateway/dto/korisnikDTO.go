package dto

import (
	"time"
)

type TipKorisnika int

const (
	CLAN TipKorisnika = iota
	BIBLIOTEKAR
	ADMIN
)

type KorisnikDTO struct {
	Email          string
	Lozinka        string
	Ime            string
	Prezime        string
	Tip            TipKorisnika
	IstekClanarine time.Time
	Sumnjiv        bool
}
