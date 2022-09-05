package dto

import "time"

type RezervacijaDTO struct {
	Id               uint
	DatumVremeIsteka time.Time
	KorisnikId       uint
	KnjigaId         uint
	Aktivno          bool
}

type RezervacijaNazivKnjigeDTO struct {
	Id               uint
	DatumVremeIsteka time.Time
	KorisnikId       uint
	KnjigaId         uint
	KnjigaNaziv      string
	Aktivno          bool
}
