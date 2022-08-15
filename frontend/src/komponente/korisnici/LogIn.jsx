import { Button, Card, CardContent, Grid, Paper, TextField } from "@mui/material";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const LogIn = () => {
    const [email, setEmail] = useState("");
    const [lozinka, setLozinka] = useState("");
    const navigation = useNavigate();




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