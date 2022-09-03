import { Button, Grid } from "@mui/material";
import React from "react";
import AddIcon from '@mui/icons-material/Add';
import KnjigeTabela from "./KnjigeTabela";
import { useState } from "react";
import KnjigaAddEditDijalog from "./KnjigaAddEditDijalog";

const KnjigeCRUDPregled = () =>{
    const [dijalogOtvoren, setDijalogOtvoren] = useState(false);

    function toggleDijalog(){
        setDijalogOtvoren(!dijalogOtvoren)
    }

    return (
        <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh' }}
        >
            <KnjigaAddEditDijalog
               otvoren={dijalogOtvoren}
               zatvoriDijalog={toggleDijalog}
               dodavanjeMod={true}
               />
            <Button variant="contained" startIcon={<AddIcon />}
                size="large"
                color = "success"
                sx={{margin: 5}}
                onClick={()=>toggleDijalog()}>
                Dodaj
            </Button>
            
            <KnjigeTabela/>
        </Grid>
      );
}

export default KnjigeCRUDPregled