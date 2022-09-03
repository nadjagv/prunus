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
import { PhotoCamera, SettingsPhone } from '@mui/icons-material';
import Putanje from '../../konstante/Putanje';
import axios from "axios";
import { useEffect } from 'react';

const KorisnikAddEditDijalog = ({otvoren, zatvoriDijalog, dodavanjeMod, korisnik, admin}) => {
  const [email, setEmail] = useState("");
  const [lozinka, setLozinka] = useState("");
  const [ime, setIme] = useState("");
  const [prezime, setPrezime] = useState("");
  const [tip, setTip] = useState(0);


  let dto = {
    Id: korisnik!=null ? korisnik.Id : 0,
    Email: email,
    Lozinka: lozinka,
    Ime: ime,
    Prezime: prezime,
    Tip: tip
  }

  useEffect(()=>{
    if (!dodavanjeMod){
      setEmail(korisnik.Email)
      setIme(korisnik.Ime)
      setPrezime(korisnik.Prezime)
      setTip(korisnik.Tip)
    }
    
    
}, [])



  function obradiPotvrdu (){
    if (dodavanjeMod){
      if (!admin){
        setTip(0)
      }
      axios
          .post(Putanje.korisniciGWURL, dto)
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
          .put(Putanje.korisniciGWURL, dto)
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
        <DialogTitle>Korisnik</DialogTitle>
        <DialogContent>
          <TextField
          disabled={!dodavanjeMod}
          margin="normal"
          label="Email"
          placeholder="Unesite email"
          fullWidth
          required
          defaultValue={email}
          onChange={(e) => {
              setEmail(e.target.value);
          }}
          ></TextField>

          {dodavanjeMod &&
          <TextField
          margin="normal"
          label="Lozinka"
          placeholder="Unesite lozinku"
          fullWidth
          required
          onChange={(e) => {
              setLozinka(e.target.value);
          }}
          ></TextField>
          }
          <TextField
          margin="normal"
          label="Ime"
          placeholder="Unesite ime"
          fullWidth
          required
          defaultValue={ime}
          onChange={(e) => {
              setIme(e.target.value);
          }}
          ></TextField>
          <TextField
          margin="normal"
          label="Prezime"
          placeholder="Unesite prezime"
          fullWidth
          required
          defaultValue={prezime}
          onChange={(e) => {
              setPrezime(e.target.value);
          }}
          ></TextField>

          {admin &&
          <Select
              id="simple-select"
              value={tip}
              label="Tip korisnika"
              fullWidth
              required
              defaultValue={tip}
              onChange={(e) => {
                console.log(e.target.value)
                setTip(e.target.value);
                
            }}
            >
              <MenuItem value={0}>Član</MenuItem>
              <MenuItem value={1}>Bibliotekar</MenuItem>
              <MenuItem value={2}>Admin</MenuItem>
            </Select>
            }
            
        </DialogContent>
        <DialogActions>
          <Button onClick={odustani}>Odustani</Button>
          <Button onClick={obradiPotvrdu}>Potvrdi</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default KorisnikAddEditDijalog