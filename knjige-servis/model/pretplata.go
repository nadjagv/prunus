package model

import "gorm.io/gorm"

type Pretplata struct {
	gorm.Model

	KorisnikId    uint   `gorm:"not null;"`
	KorisnikEmail string `gorm:"not null;"`
	KnjigaId      uint   `gorm:"not null;"`
}

func (Pretplata) TableName() string {
	return "pretplate"
}

type PretplataDTO struct {
	Id uint

	KorisnikId    uint
	KorisnikEmail string
	KnjigaId      uint
}

func (pretplata *Pretplata) MapirajNaDTO() PretplataDTO {
	return PretplataDTO{
		Id:            pretplata.ID,
		KorisnikId:    pretplata.KorisnikId,
		KorisnikEmail: pretplata.KorisnikEmail,
		KnjigaId:      pretplata.KnjigaId,
	}
}

func (pretplata *PretplataDTO) MapirajNaObjekat() Pretplata {
	return Pretplata{
		KorisnikId:    pretplata.KorisnikId,
		KorisnikEmail: pretplata.KorisnikEmail,
		KnjigaId:      pretplata.KnjigaId,
	}
}
