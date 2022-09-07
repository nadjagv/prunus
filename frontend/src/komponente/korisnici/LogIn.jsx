import { Button, Card, CardContent, FormControl, Grid, IconButton, InputAdornment, InputLabel, OutlinedInput, Paper, TextField } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import AuthServis from "../../servisi/AuthServis";
import axios from "axios";
import Putanje from "../../konstante/Putanje";
import { PropaneSharp, Visibility, VisibilityOff } from "@mui/icons-material";

const LogIn = ({handleUlogovan}) => {
    const [email, setEmail] = useState("");
    const [lozinka, setLozinka] = useState("");
    const navigation = useNavigate();
    const [prikaziLozinku, setPrikaziLozinku] = useState("");

    const handlePrikaziLozinku = () => {
        setPrikaziLozinku(!prikaziLozinku)
      };
    
    const handleMouseDownLozinka = (event) => {
    event.preventDefault();
    };

    useEffect(() => {
        if (AuthServis.preuzmiKorisnika()) {
        //   axios.post(environment.baseURL + "auth/logout").then((response) => {
        //     AuthService.removeUser();
        //   });
        }
      }, []);
    
      let kredencijali = {
        Email: email,
        Lozinka: lozinka,
      };
      const ulogujKorisnika = () => {
        axios
          .post(Putanje.korisniciGWURL + "/login", kredencijali)
          .then((response) => {
            console.log(response.data);
            AuthServis.postaviKorisnika(response.data);
            handleUlogovan(true);
            navigation("/");
          })
          .catch((error) => {
            alert("Loši kredencijali.");
          });

          
      };


    return (
        <div className="container">
        <Card>
            <CardContent style={{justifyContent: 'center', alignItems: 'center'}}>
                <TextField
                margin="normal"
                label="Email"
                placeholder="Unesite Vaš email"
                fullWidth
                required
                onChange={(e) => {
                    setEmail(e.target.value);
                }}
                ></TextField>
                <FormControl variant="outlined" fullWidth>
                <InputLabel htmlFor="lozinka">Lozinka</InputLabel>
                <OutlinedInput
                        id="lozinka"
                        label="Lozinka"
                        placeholder="Unesite Vašu lozinku"
                        fullWidth
                        required
                        type={prikaziLozinku ? 'text' : 'password'}
                        onChange={(e) => {
                            setLozinka(e.target.value);
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
                <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center', marginTop: 10}}>
                    <Button 
                    type="submit"
                    color="primary"
                    variant="contained"
                    onClick={() => {
                        ulogujKorisnika()
                    }}
                    >
                    Log In
                    </Button>
                </div>
            </CardContent>
        </Card>
        
    </div>
  );

}

export default LogIn