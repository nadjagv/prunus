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

const IznajmljivanjeAddDijalog = ({otvoren, zatvoriDijalog, knjigaId}) => {
  const [email, setEmail] = useState("");

  let dto = {
    Email: email,
    KnjigaId: knjigaId
  }

  function obradiPotvrdu (){
    axios
          .post(Putanje.iznajmljivanjaGWURL, dto)
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
        <DialogTitle>Iznajmljivanje</DialogTitle>
        <DialogContent>
          <TextField
          margin="normal"
          label="Email"
          placeholder="Unesite email korisnika"
          fullWidth
          required
          onChange={(e) => {
              setEmail(e.target.value);
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

export default IznajmljivanjeAddDijalog