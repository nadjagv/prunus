package dto

type StatusRecenzije int

const (
	KREIRANO StatusRecenzije = iota
	ODOBRENO
	ODBIJENO
)

type RecenzijaDTO struct {
	Id          uint
	Korisnik_id uint
	Knjiga_id   uint
	Ocena       uint
	Komentar    string
	Obrisano    bool
	Status      string
}

type RecenzijaNazivEmailDTO struct {
	Id            uint
	KorisnikId    uint
	KnjigaId      uint
	Ocena         uint
	Komentar      string
	Obrisano      bool
	Status        string
	KorisnikEmail string
	KnjigaNaziv   string
}
