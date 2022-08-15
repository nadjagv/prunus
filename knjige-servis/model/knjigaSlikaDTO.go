package model

import (
	"fmt"
	"image"

	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "image/jpeg"
)

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

func (knjiga *Knjiga) MapirajNaSlikaDTO() KnjigaSlikaDTO {
	slikaBase64, err := konvertujUBase64(knjiga.Slika)
	if err != nil {
		fmt.Println(err)
		return KnjigaSlikaDTO{}
	}

	return KnjigaSlikaDTO{
		Id:               knjiga.ID,
		Isbn:             knjiga.Isbn,
		Naziv:            knjiga.Naziv,
		ImeAutora:        knjiga.ImeAutora,
		PrezimeAutora:    knjiga.PrezimeAutora,
		Opis:             knjiga.Opis,
		Zanr:             knjiga.Zanr,
		BrojStrana:       knjiga.BrojStrana,
		GodinaNastanka:   knjiga.GodinaNastanka,
		UkupnaKolicina:   knjiga.UkupnaKolicina,
		TrenutnoDostupno: knjiga.TrenutnoDostupno,
		ProsecnaOcena:    knjiga.ProsecnaOcena,
		Slika:            slikaBase64,
	}
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func konvertujUBase64(putanja string) (string, error) {
	// Read the entire file into a byte slice
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	sep := string(os.PathSeparator)

	bytes, err := ioutil.ReadFile(dir + sep + "slike" + sep + putanja)
	if err != nil {
		return "", err
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	return base64Encoding, nil
}

func ucitajSlikuSaPutanje(putanja string) (image.Image, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	sep := string(os.PathSeparator)

	f, err := os.Open(dir + sep + "slike" + sep + putanja)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	f.Seek(0, 0)

	slika, _, err := image.Decode(f)
	return slika, err
}
