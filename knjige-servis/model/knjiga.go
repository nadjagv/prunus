package model

import (
	"time"

	"gorm.io/gorm"
)

type ZanrEnum int

const (
	NaucnaFantastika     ZanrEnum = iota //0
	Ljubavni                             //1
	Horor                                //2
	Triler                               //3
	Avantura                             //4
	Biografija                           //5
	PopularnaPsihologija                 //6
	OpstaInteresovanja                   //7
	StrucnaLiteratura                    //8
	StraniJezik                          //9
	Poezija                              //10
	Decije                               //11
	Ostalo                               //12
)

type Knjiga struct {
	gorm.Model

	Isbn             string    `gorm:"not null;"`
	Naziv            string    `gorm:"not null;"`
	ImeAutora        string    `gorm:"not null;"`
	PrezimeAutora    string    `gorm:"not null;"`
	Opis             string    `gorm:"not null;"`
	Zanr             ZanrEnum  `gorm:"not null;default:12"`
	BrojStrana       uint      `gorm:"not null;"`
	GodinaNastanka   time.Time `gorm:"not null;"`
	UkupnaKolicina   uint      `gorm:"not null;"`
	TrenutnoDostupno uint      `gorm:"not null;"`
	ProsecnaOcena    float64   `gorm:"not null;"`
	Slika            string
}

type Tabler interface {
	TableName() string
}

func (Knjiga) TableName() string {
	return "knjige"
}
