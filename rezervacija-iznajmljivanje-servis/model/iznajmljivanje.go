package model

import (
	"time"

	"gorm.io/gorm"
)

type Iznajmljivanje struct {
	gorm.Model

	DatumVremeIznajmljivanja time.Time `gorm:"not null;"`
	RokVracanja              time.Time `gorm:"not null;"`
	DatumVremeVracanja       time.Time
	ZakasneloVracanje        bool `gorm:"not null;"`
	KorisnikId               uint `gorm:"not null;"`
	KnjigaId                 uint `gorm:"not null;"`
	Produzeno                bool `gorm:"not null;"`
	Aktivno                  bool `gorm:"not null;"`
}

type Tabler interface {
	TableName() string
}

func (Iznajmljivanje) TableName() string {
	return "iznajmljivanja"
}
