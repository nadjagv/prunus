package model

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
