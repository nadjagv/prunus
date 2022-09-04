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

const ObrazlozenjeBlokiranjaDijalog = ({otvoren, zatvoriDijalog, korisnikId}) => {
  const [obrazlozenje, setObrazlozenje] = useState("");



  function obradiPotvrdu (){
    axios
    .put(`${Putanje.korisniciGWURL}/blokiraj/${korisnikId}`, obrazlozenje)
    .then((response) => {
      console.log(response.data);
      zatvoriDijalog()
      alert("Korisnik blokiran.");
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
        <DialogTitle>Korisnik</DialogTitle>
        <DialogContent>
          <TextField
          margin="normal"
          label="Obrazloženje"
          placeholder="Unesite obrarzloženje"
          multiline
          fullWidth
          required
          onChange={(e) => {
              setObrazlozenje(e.target.value);
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

export default ObrazlozenjeBlokiranjaDijalog