package model

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

type NovoIznajmljivanjeDTO struct {
	Email    string
	KnjigaId uint
}

func (iznajmljivanje *Iznajmljivanje) MapirajNaDTO() IznajmljivanjeDTO {
	return IznajmljivanjeDTO{
		Id:                       iznajmljivanje.ID,
		DatumVremeIznajmljivanja: iznajmljivanje.DatumVremeIznajmljivanja,
		RokVracanja:              iznajmljivanje.RokVracanja,
		DatumVremeVracanja:       iznajmljivanje.DatumVremeVracanja,
		ZakasneloVracanje:        iznajmljivanje.ZakasneloVracanje,
		KorisnikId:               iznajmljivanje.KorisnikId,
		KnjigaId:                 iznajmljivanje.KnjigaId,
		Produzeno:                iznajmljivanje.Produzeno,
		Aktivno:                  iznajmljivanje.Aktivno,
	}
}

func (iznajmljivanje *IznajmljivanjeDTO) MapirajNaObjekat() Iznajmljivanje {
	return Iznajmljivanje{
		DatumVremeIznajmljivanja: iznajmljivanje.DatumVremeIznajmljivanja,
		RokVracanja:              iznajmljivanje.RokVracanja,
		DatumVremeVracanja:       iznajmljivanje.DatumVremeVracanja,
		ZakasneloVracanje:        iznajmljivanje.ZakasneloVracanje,
		KorisnikId:               iznajmljivanje.KorisnikId,
		KnjigaId:                 iznajmljivanje.KnjigaId,
		Produzeno:                iznajmljivanje.Produzeno,
		Aktivno:                  iznajmljivanje.Aktivno,
	}
}
