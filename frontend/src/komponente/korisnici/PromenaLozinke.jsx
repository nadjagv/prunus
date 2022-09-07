import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { useState } from 'react';
import { Card, CardContent, FormControl, IconButton, InputAdornment, InputLabel, MenuItem, OutlinedInput, Select } from '@mui/material';
import { PhotoCamera, SettingsPhone, Visibility, VisibilityOff } from '@mui/icons-material';
import Putanje from '../../konstante/Putanje';
import axios from "axios";
import { useEffect } from 'react';
import AuthServis from '../../servisi/AuthServis';
import { useNavigate } from 'react-router-dom';
import { Stack } from '@mui/system';


const PromenaLozinke = () => {
    const [stara, setStara] = useState("");
    const [nova, setNova] = useState("");
    const [prikaziLozinku, setPrikaziLozinku] = useState("");
    const [prikaziLozinkuNova, setPrikaziLozinkuNova] = useState("");

    const navigate = useNavigate()

    const handlePrikaziLozinku = () => {
        setPrikaziLozinku(!prikaziLozinku)
      };
    
    const handleMouseDownLozinka = (event) => {
    event.preventDefault();
    };

    const handlePrikaziLozinkuNova = () => {
        setPrikaziLozinkuNova(!prikaziLozinkuNova)
      };
    
    const handleMouseDownLozinkaNova = (event) => {
    event.preventDefault();
    };
  
  
    let dto = {
      Stara: stara,
      Nova: nova,
      KorisnikId: AuthServis.preuzmiKorisnika().Id,
    }


  function obradiPotvrdu (){
    axios
          .put(Putanje.korisniciGWURL+"/lozinka", dto)
          .then((response) => {
            console.log(response.data);
            alert("Zahtev uspešno obrađen!");
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
  }

  return (
    <div className='container'>
      <Card>
        <CardContent >

        <Stack spacing={2} direction="column">
        <FormControl  variant="outlined">
            <InputLabel htmlFor="stara-lozinka">Stara lozinka</InputLabel>
            <OutlinedInput
            id="stara-lozinka"
            label="Stara lozinka"
            placeholder="Unesite staru lozinku"
            fullWidth
            required
            type={prikaziLozinku ? 'text' : 'password'}
            onChange={(e) => {
            setStara(e.target.value);
            }}
            endAdornment={
            <InputAdornment position="end">
                <IconButton
                aria-label="toggle password visibility"
                onClick={handlePrikaziLozinku}
                onMouseDown={handleMouseDownLozinka}
                edge="end"
                >
                {prikaziLozinku ? <VisibilityOff /> : <Visibility />}
                </IconButton>
            </InputAdornment>
            }
            ></OutlinedInput>
        </FormControl>
        <FormControl variant="outlined">
            <InputLabel htmlFor="nova-lozinka">Nova lozinka</InputLabel>
            <OutlinedInput
                    id="nova-lozinka"
                    label="Nova lozinka"
                    placeholder="Unesite novu lozinku"
                    fullWidth
                    required
                    type={prikaziLozinkuNova ? 'text' : 'password'}
                    onChange={(e) => {
                        setNova(e.target.value);
                    }}
                    endAdornment={
                        <InputAdornment position="end">
                        <IconButton
                            aria-label="toggle password visibility"
                            onClick={handlePrikaziLozinkuNova}
                            onMouseDown={handleMouseDownLozinkaNova}
                            edge="end"
                        >
                            {prikaziLozinkuNova ? <VisibilityOff /> : <Visibility />}
                        </IconButton>
                        </InputAdornment>
                    }
                ></OutlinedInput>
            </FormControl>
            </Stack>
            <Stack spacing={2} direction="row" style={{display: 'flex', justifyContent: 'center', alignItems: 'center', marginTop: 10}}>
            
              <Button 
                color="primary"
                variant="contained"
                onClick={() => {
                    navigate("/knjige")
                }}
                >
                Odustani
                </Button>
                
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
              
                
            </Stack>
            
        </CardContent>
      </Card>
    </div>
  );
}

export default PromenaLozinke