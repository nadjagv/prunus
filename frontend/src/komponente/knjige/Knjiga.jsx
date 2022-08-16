import { TryRounded } from "@mui/icons-material";
import { Box, Card, CardContent, Grid } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Zanr from "../../enumeracije/Zanr"
import Putanje from "../../konstante/Putanje";

const Knjiga = () =>{
    const { id } = useParams()

    const [knjiga, setKnjiga] = useState([])

    useEffect(()=>{
        preuzmiPoId()
        
    }, [])

    const preuzmiPoId = async () => {
        console.log(id)
        const response = await fetch (`${Putanje.knjigeGWURL}/${id}`)
        const data = await response.json();
        console.log(data)
        setKnjiga(data)
    }


    return (
        <Box className="container">
          <Grid item container spacing={50} columns={10} zeroMinWidth>
            <Grid xs={3} item={true} zeroMinWidth sx={{ marginLeft: 10 }}>
                <div className="jednaKnjigaSlika">
                    <img src={knjiga.Slika} alt={knjiga.Naziv} />
                </div>
                
                
            </Grid>
            <Grid xs={5} item={true} zeroMinWidth sx={{marginRight: 10, width: 1/2}}>
                <Card sx={{marginTop: 3}}>
                    <CardContent>
                        <h3>{knjiga.ImeAutora} {knjiga.PrezimeAutora}</h3>
                        <h2 style={{color: "#e07421"}}>{knjiga.Naziv}</h2>
                        <h4>ISBN: {knjiga.Isbn}</h4>
                        <h4>Zanr: {Object.keys(Zanr).find(key => Zanr[key] === knjiga.Zanr)}</h4>

                        <br/>
                        <p>{knjiga.Opis}</p>

                        <br/>
                        <p> Broj strana: {knjiga.BrojStrana}</p>
                        <p> Godina nastanka: {knjiga.GodinaNastanka}</p>
                        <p> Ocena: {knjiga.ProsecnaOcena}</p>


                    </CardContent>
                </Card>
            </Grid>
          </Grid>
        </Box>
      );
}

export default Knjiga