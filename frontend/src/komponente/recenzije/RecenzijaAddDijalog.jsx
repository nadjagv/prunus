import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { useState } from 'react';
import { IconButton, MenuItem, Rating, Select } from '@mui/material';
import { PhotoCamera, SettingsPhone } from '@mui/icons-material';
import Putanje from '../../konstante/Putanje';
import axios from "axios";
import { useEffect } from 'react';
import AuthServis from '../../servisi/AuthServis';

const RecenzijaAddDijalog = ({otvoren, zatvoriDijalog, knjigaId}) => {
  const [komentar, setKomentar] = useState("");
  const [ocena, setOcena] = useState(0);

  let dto ={
    id: 0,
    knjiga_id: knjigaId,
    komentar: komentar,
    korisnik_id:AuthServis.preuzmiKorisnika().Id,
    obrisano: false,
    ocena: parseInt(ocena),
    status: "KREIRANO"
    }



  function obradiPotvrdu (){
    axios
          .post(Putanje.recenzijeGWURL, dto)
          .then((response) => {
            console.log(response.data);
            zatvoriDijalog(true)
            alert("Zahtev uspešno obrađen!");
            
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
  }
  function odustani(){
    zatvoriDijalog(false)
  }

  return (
    <div>
      <Dialog open={otvoren} onClose={zatvoriDijalog}>
        <DialogTitle>Recenzija</DialogTitle>
        <DialogContent>

        <Rating
            name="simple-controlled"
            value={parseInt(ocena)}
            onChange={(e) => {
                setOcena(e.target.value)
            }}
        />
          <TextField
          margin="normal"
          multiline
          label="Komentar"
          placeholder="Unesite komentar"
          fullWidth
          required
          onChange={(e) => {
              setKomentar(e.target.value)
          }}
          ></TextField>

          
            
        </DialogContent>
        <DialogActions>
          <Button onClick={odustani}>Odustani</Button>
          <Button onClick={obradiPotvrdu}>Potvrdi</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default RecenzijaAddDijalog