import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { useState } from 'react';
import { IconButton, MenuItem, Select } from '@mui/material';
import { PhotoCamera } from '@mui/icons-material';
import Putanje from '../../konstante/Putanje';
import axios from "axios";
import { useEffect } from 'react';

const KnjigaAddEditDijalog = ({otvoren, zatvoriDijalog, dodavanjeMod, knjiga}) => {
  const [naziv, setNaziv] = useState("");
  const [isbn, setIsbn] = useState("");
  const [imeAutora, setImeAutora] = useState("");
  const [prezimeAutora, setPrezimeAutora] = useState("");
  const [opis, setOpis] = useState("");
  const [zanr, setZanr] = useState(0);
  const [brojStrana, setBrojStrana] = useState(0);
  const [godinaNastanka, setGodinaNastanka] = useState(1900);
  const [ukupnaKolicina, setUkupnaKolicina] = useState(1);

  const [base64URL, setBase64URL] = useState("");


  let dto = {
    Id: knjiga!=null ? knjiga.Id : 0,
    Naziv: naziv,
    Isbn: isbn,
    ImeAutora: imeAutora,
    PrezimeAutora: prezimeAutora,
    Opis: opis,
    Zanr: zanr,
    BrojStrana: parseInt(brojStrana),
    GodinaNastanka: parseInt(godinaNastanka),
    UkupnaKolicina: parseInt(ukupnaKolicina),
    Slika: base64URL
  }

  useEffect(()=>{
    if (!dodavanjeMod){
      setNaziv(knjiga.Naziv)
      setIsbn(knjiga.Isbn)
      setImeAutora(knjiga.ImeAutora)
      setPrezimeAutora(knjiga.PrezimeAutora)
      setOpis(knjiga.Opis)
      setZanr(knjiga.Zanr)
      setBrojStrana(knjiga.BrojStrana)
      setGodinaNastanka(knjiga.GodinaNastanka)
      setUkupnaKolicina(knjiga.UkupnaKolicina)
    }
    
    
}, [])


  const handleFileInputChange = e => {
    let reader = new FileReader();
    reader.readAsDataURL(e.target.files[0]);
      reader.onload = () => {
        setBase64URL(reader.result)
      };
  };

  function obradiPotvrdu (){
    if (dodavanjeMod){
      axios
          .post(Putanje.knjigeGWURL, dto)
          .then((response) => {
            console.log(response.data);
            zatvoriDijalog(true)
            alert("Zahtev uspešno obrađen!");
            
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }else{
      axios
          .put(Putanje.knjigeGWURL, dto)
          .then((response) => {
            console.log(response.data);
            zatvoriDijalog(true)
            alert("Zahtev uspešno obrađen!");
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }
  }
  function odustani(){
    zatvoriDijalog(false)
  }

  return (
    <div>
      <Dialog open={otvoren} onClose={zatvoriDijalog}>
        <DialogTitle>Knjiga</DialogTitle>
        <DialogContent>
          <TextField
          margin="normal"
          label="Naziv"
          placeholder="Unesite naziv knjige"
          fullWidth
          required
          defaultValue={naziv}
          onChange={(e) => {
              setNaziv(e.target.value);
          }}
          ></TextField>

          <TextField
          margin="normal"
          label="ISBN"
          placeholder="Unesite ISBN"
          disabled={!dodavanjeMod}
          fullWidth
          required
          defaultValue={isbn}
          onChange={(e) => {
              setIsbn(e.target.value);
          }}
          ></TextField>
          <TextField
          margin="normal"
          label="Ime autora"
          placeholder="Unesite ime autora"
          fullWidth
          required
          defaultValue={imeAutora}
          onChange={(e) => {
              setImeAutora(e.target.value);
          }}
          ></TextField>
          <TextField
          margin="normal"
          label="Prezime autora"
          placeholder="Unesite prezime autora"
          fullWidth
          required
          defaultValue={prezimeAutora}
          onChange={(e) => {
              setPrezimeAutora(e.target.value);
          }}
          ></TextField>
          <TextField
          margin="normal"
          label="Opis"
          placeholder="Unesite opis"
          fullWidth
          required
          defaultValue={opis}
          onChange={(e) => {
              setOpis(e.target.value);
          }}
          ></TextField>

          <Select
              id="simple-select"
              value={zanr}
              label="Žanr"
              fullWidth
              required
              defaultValue={zanr}
              onChange={(e) => {
                console.log(e.target.value)
                setZanr(e.target.value);
                
            }}
            >
              <MenuItem value={0}>Naučna fantastika</MenuItem>
              <MenuItem value={1}>Ljubavni</MenuItem>
              <MenuItem value={2}>Klasik</MenuItem>
              <MenuItem value={3}>Horor</MenuItem>
              <MenuItem value={4}>Triler</MenuItem>
              <MenuItem value={5}>Avantura</MenuItem>
              <MenuItem value={6}>Biografija</MenuItem>
              <MenuItem value={7}>Popularna psihologija</MenuItem>
              <MenuItem value={8}>Opšta interesovanja</MenuItem>
              <MenuItem value={9}>Stručna literatura</MenuItem>
              <MenuItem value={10}>Strani jezik</MenuItem>
              <MenuItem value={11}>Poezija</MenuItem>
              <MenuItem value={12}>Dečije</MenuItem>
              <MenuItem value={13}>Ostalo</MenuItem>
            </Select>

            <TextField
            margin="normal"
            label="Broj strana"
            type="number"
            defaultValue={brojStrana}
            placeholder="Unesite broj strana"
            fullWidth
            required
            onChange={(e) => {
                setBrojStrana(e.target.value);
            }}
            ></TextField>

            <TextField
            margin="normal"
            label="Godina nastanka"
            defaultValue={godinaNastanka}
            type="number"
            placeholder="Unesite godinu nastanka"
            fullWidth
            required
            onChange={(e) => {
                setGodinaNastanka(e.target.value);
            }}
            ></TextField>

            <TextField
            margin="normal"
            label="Ukupna količina"
            defaultValue={ukupnaKolicina}
            type="number"
            placeholder="Unesite ukupnu količinu"
            fullWidth
            required
            onChange={(e) => {
                setUkupnaKolicina(e.target.value);
            }}
            ></TextField>


          <IconButton color="secondary" aria-label="upload picture" component="label">
            <input hidden type="file" name="file" onChange={handleFileInputChange} />
            <PhotoCamera />
          </IconButton>
        </DialogContent>
        <DialogActions>
          <Button onClick={odustani}>Odustani</Button>
          <Button onClick={obradiPotvrdu}>Potvrdi</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default KnjigaAddEditDijalog