package model

import (
	"time"

	"gorm.io/gorm"
)

type TipKorisnika int

const (
	CLAN TipKorisnika = iota
	BIBLIOTEKAR
	ADMIN
)

type Korisnik struct {
	gorm.Model
	Email          string       `gorm:"not null;"`
	Lozinka        string       `gorm:"not null;size:255"`
	Ime            string       `gorm:"not null;"`
	Prezime        string       `gorm:"not null;"`
	Tip            TipKorisnika `gorm:"not null;"`
	IstekClanarine time.Time
	Sumnjiv        bool `gorm:"not null;default:false"`
	Blokiran       bool `gorm:"not null;default:false"`
}

type Tabler interface {
	TableName() string
}

func (Korisnik) TableName() string {
	return "korisnici"
}
