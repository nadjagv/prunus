package model

import (
	"time"

	"gorm.io/gorm"
)

type Rezervacija struct {
	gorm.Model

	DatumVremeIsteka time.Time `gorm:"not null;"`
	KorisnikId       uint      `gorm:"not null;"`
	KnjigaId         uint      `gorm:"not null;"`
	Aktivno          bool      `gorm:"not null;"`
}

func (Rezervacija) TableName() string {
	return "rezervacije"
}
