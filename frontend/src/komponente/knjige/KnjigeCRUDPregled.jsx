import { Button, Grid } from "@mui/material";
import React from "react";
import AddIcon from '@mui/icons-material/Add';
import KnjigeTabela from "./KnjigeTabela";

const KnjigeCRUDPregled = ({knjiga}) =>{
    return (
        <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh' }}
        >
            <Button variant="contained" startIcon={<AddIcon />}
                size="large"
                color = "success"
                sx={{margin: 5}}
                onClick={() => alert("dodaj")}>
                Dodaj
            </Button>
            <KnjigeTabela/>
        </Grid>
      );
}

export default KnjigeCRUDPregled