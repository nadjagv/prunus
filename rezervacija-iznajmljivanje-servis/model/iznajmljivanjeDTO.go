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
	}
}
