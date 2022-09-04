import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { useState } from 'react';
import { Card, CardContent, IconButton, MenuItem, Select } from '@mui/material';
import { PhotoCamera, SettingsPhone } from '@mui/icons-material';
import Putanje from '../../konstante/Putanje';
import axios from "axios";
import { useEffect } from 'react';
import AuthServis from '../../servisi/AuthServis';
import { useNavigate } from 'react-router-dom';

const KorisnikNalogForma = ({ dodavanjeMod}) => {
    const [email, setEmail] = useState("");
    const [lozinka, setLozinka] = useState("");
    const [ime, setIme] = useState("");
    const [prezime, setPrezime] = useState("");
    const [tip, setTip] = useState(0);
    const [id, setId] = useState(0);
    const [korisnik, setKorisnik] = useState({});
    const navigation = useNavigate();
  
  
    let dto = {
      Id: id,
      Email: email,
      Lozinka: lozinka,
      Ime: ime,
      Prezime: prezime,
      Tip: tip
    }
  
    useEffect(()=>{
      if (!dodavanjeMod){
        preuzmiKorisnika()
      }}, [])

    const preuzmiKorisnika = () => {
        let ulogovanEmail = AuthServis.preuzmiKorisnika().Email
        axios
        .get(`${Putanje.korisniciGWURL}/email/${ulogovanEmail}`)
        .then((response) => {
            console.log(response.data)
            setId(response.data.Id)
            setEmail(response.data.Email)
            setIme(response.data.Ime)
            setPrezime(response.data.Prezime)
            setTip(response.data.Tip)
            
        })
        .catch((error) => {
            alert("Neuspešno dobavljanje korisnika.");
        });

        
    }



  function obradiPotvrdu (){
    console.log(email + "yheheh")
    if (dodavanjeMod){
        setTip(0)
      axios
          .post(Putanje.korisniciGWURL, dto)
          .then((response) => {
            console.log(response.data);
            alert("Zahtev uspešno obrađen!");
            navigation("/login")
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }else{
      axios
          .put(Putanje.korisniciGWURL, dto)
          .then((response) => {
            console.log(response.data);
            alert("Zahtev uspešno obrađen!");
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }
  }

  return (
    <div className='container'>
      <Card>
        <CardContent>
          <TextField
          disabled={!dodavanjeMod}
          margin="normal"
          label="Email"
          placeholder="Unesite email"
          fullWidth
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
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
          value={ime}
          onChange={(e) => setIme(e.target.value)}
          ></TextField>
          <TextField
          margin="normal"
          label="Prezime"
          placeholder="Unesite prezime"
          fullWidth
          required
          value={prezime}
          onChange={(e) => {
            setPrezime(e.target.value);
          }}
          ></TextField>


            <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center', marginTop: 10}}>
                    <Button 
                    type="submit"
                    color="primary"
                    variant="contained"
                    onClick={() => {
                        obradiPotvrdu()
                    }}
                    >
                    Potvrdi
                    </Button>
                </div>
            
        </CardContent>
      </Card>
    </div>
  );
}

export default KorisnikNalogForma