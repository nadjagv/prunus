package dto

import "time"

type IznajmljivanjeDTO struct {
	Id                       uint
	DatumVremeIznajmljivanja time.Time
	RokVracanja              time.Time
	DatumVremeVracanja       time.Time
	ZakasneloVracanje        bool
	KorisnikId               uint
	KnjigaId                 uint
	Produzeno                bool
	Aktivno                  bool
}
