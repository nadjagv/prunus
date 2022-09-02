import { Button, Card, CardContent, Grid, Paper, TextField } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import AuthServis from "../../servisi/AuthServis";
import axios from "axios";
import Putanje from "../../konstante/Putanje";
import { PropaneSharp } from "@mui/icons-material";

const LogIn = ({handleUlogovan}) => {
    const [email, setEmail] = useState("");
    const [lozinka, setLozinka] = useState("");
    const navigation = useNavigate();

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
                <TextField
                label="Lozinka"
                placeholder="Unesite lozinku"
                type="password"
                onChange={(e) => {
                    setLozinka(e.target.value);
                }}
                fullWidth
                required
                ></TextField>
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