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
