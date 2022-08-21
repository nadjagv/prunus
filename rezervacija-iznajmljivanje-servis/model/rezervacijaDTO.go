package model

import "time"

type RezervacijaDTO struct {
	Id               uint
	DatumVremeIsteka time.Time
	KorisnikId       uint
	KnjigaId         uint
	Aktivno          bool
}

func (rezervacija *Rezervacija) MapirajNaDTO() RezervacijaDTO {
	return RezervacijaDTO{
		Id:               rezervacija.ID,
		DatumVremeIsteka: rezervacija.DatumVremeIsteka,
		KorisnikId:       rezervacija.KorisnikId,
		KnjigaId:         rezervacija.KnjigaId,
		Aktivno:          rezervacija.Aktivno,
	}
}

func (rezervacija *RezervacijaDTO) MapirajNaObjekat() Rezervacija {
	return Rezervacija{
		DatumVremeIsteka: rezervacija.DatumVremeIsteka,
		KorisnikId:       rezervacija.KorisnikId,
		KnjigaId:         rezervacija.KnjigaId,
		Aktivno:          rezervacija.Aktivno,
	}
}