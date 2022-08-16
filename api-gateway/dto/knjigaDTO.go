package dto

type ZanrEnum int

const (
	NaucnaFantastika     ZanrEnum = iota //0
	Ljubavni                             //1
	Klasik                               //2
	Horor                                //3
	Triler                               //4
	Avantura                             //5
	Biografija                           //6
	PopularnaPsihologija                 //7
	OpstaInteresovanja                   //8
	StrucnaLiteratura                    //9
	StraniJezik                          //10
	Poezija                              //11
	Decije                               //12
	Ostalo                               //13
)

type KnjigaDTO struct {
	Id               uint
	Isbn             string   `gorm:"not null;"`
	Naziv            string   `gorm:"not null;"`
	ImeAutora        string   `gorm:"not null;"`
	PrezimeAutora    string   `gorm:"not null;"`
	Opis             string   `gorm:"not null;"`
	Zanr             ZanrEnum `gorm:"not null;default:12"`
	BrojStrana       uint     `gorm:"not null;"`
	GodinaNastanka   uint     `gorm:"not null;"`
	UkupnaKolicina   uint     `gorm:"not null;"`
	TrenutnoDostupno uint     `gorm:"not null;"`
	ProsecnaOcena    float64  `gorm:"not null;"`
}

type KnjigaSlikaDTO struct {
	Id               uint
	Isbn             string
	Naziv            string
	ImeAutora        string
	PrezimeAutora    string
	Opis             string
	Zanr             ZanrEnum
	BrojStrana       uint
	GodinaNastanka   uint
	UkupnaKolicina   uint
	TrenutnoDostupno uint
	ProsecnaOcena    float64
	Slika            string
}
