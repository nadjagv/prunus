package dto

type IzvestajDTO struct {
	Broj_iznajmljivanja     uint
	Broj_zakasnela_vracanja uint
	Broj_produzenja         uint

	Knjiga1 uint
	Knjiga2 uint
	Knjiga3 uint

	Broj_korisnika  uint
	Broj_sumnjivih  uint
	Broj_blokiranih uint
}
