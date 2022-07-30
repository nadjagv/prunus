package model

import (
	"time"

	"gorm.io/gorm"
)

type Iznajmljivanje struct {
	gorm.Model

	DatumVremeIznajmljivanja time.Time `gorm:"not null;"`
	RokVracanja              time.Time `gorm:"not null;"`
	DatumVremeVracanja       time.Time `gorm:"not null;"`
	ZakasneloVracanje        bool      `gorm:"not null;"`
	KorisnikId               uint      `gorm:"not null;"`
	KnjigaId                 uint      `gorm:"not null;"`
}

type Tabler interface {
	TableName() string
}

func (Iznajmljivanje) TableName() string {
	return "iznajmljivanja"
}
