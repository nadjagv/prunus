import { Button, Card, CardContent, Grid, Stack, TextField, Typography } from "@mui/material";
import { DesktopDatePicker, LocalizationProvider } from "@mui/x-date-pickers";
import React, { useState } from "react";
import Putanje from "../../konstante/Putanje";
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import axios from "axios";
import { useEffect } from "react";
import KnjigaKartica from "../knjige/KnjigaKartica";

const IzvestajPregled = () => {
    const [izvestaj, setIzvestaj] = useState(null);
    const [pocetak, setPocetak] = useState(0);
    const [kraj, setKraj] = useState(0);
    const [knjiga1, setKnjiga1] = useState(null);
    const [knjiga2, setKnjiga2] = useState(null);
    const [knjiga3, setKnjiga3] = useState(null);
    useEffect(()=>{

        setPocetak(new Date())
        setKraj(new Date())
        
    }, [])

    useEffect(()=>{
        preuzmiKnjigePoId()   
    }, [izvestaj])

    function preuzmiIzvestaj(){
        axios
          .get(`${Putanje.izvestajiGWURL}`, {params: {pocetak:pocetak.getTime(), kraj:kraj.getTime()}})
          .then((response) => {
            setIzvestaj(response.data)
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }

    const preuzmiKnjigePoId =  () => {
        if (izvestaj==null){
            return
        }

        if (izvestaj.Knjiga1 != 0){
            axios
            .get(`${Putanje.knjigeGWURL}/${izvestaj.Knjiga1}`)
            .then((response) => {
                console.log(response.data)
                setKnjiga1(response.data) 
            })
            .catch((error) => {
                console.log("Neuspešno dobavljanje knjige.");
            });
        }

        if (izvestaj.Knjiga2 != 0){
            axios
            .get(`${Putanje.knjigeGWURL}/${izvestaj.Knjiga2}`)
            .then((response) => {
                console.log(response.data)
                setKnjiga2(response.data) 
            })
            .catch((error) => {
                console.log("Neuspešno dobavljanje knjige.");
            });
        }

        if (izvestaj.Knjiga3 != 0){
            axios
            .get(`${Putanje.knjigeGWURL}/${izvestaj.Knjiga3}`)
            .then((response) => {
                console.log(response.data)
                setKnjiga3(response.data) 
            })
            .catch((error) => {
                console.log("Neuspešno dobavljanje knjige.");
            });
        }
        

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
            <Stack spacing={2} direction="row" sx={{margin: 5}}>
            <LocalizationProvider dateAdapter={AdapterDateFns}>
            <DesktopDatePicker
                label="Od"
                inputFormat="dd/MM/yyyy"
                value={pocetak}
                onChange={(e) => setPocetak(e)}
                renderInput={(params) => <TextField {...params} />}
            />

            <DesktopDatePicker
                label="Do"
                inputFormat="dd/MM/yyyy"
                minDate={pocetak}
                value={kraj}
                onChange={(e) => setKraj(e)}
                renderInput={(params) => <TextField {...params} />}
            />
            </LocalizationProvider>

            <Button 
                color="primary" 
                variant="contained"
                onClick={() => preuzmiIzvestaj()}>
                    Preuzmi izveštaj
            </Button>
            </Stack>
            <br></br>
            {izvestaj!=null &&
            <Grid xs={5} item={true} zeroMinWidth sx={{marginRight: 10, width: 3/4}}>
                <Card>
                    <CardContent>
                    <Typography variant="h5" gutterBottom component="div">
                            <b>Izveštaj</b>
                        </Typography>
                        <Typography variant="h6" gutterBottom component="div">
                            <b>Iznajmljivanja</b>
                        </Typography>
                        <Typography  gutterBottom component="div">
                            Broj iznajmljivanja: {izvestaj.Broj_iznajmljivanja}<br/>
                            Broj produženja: {izvestaj.Broj_produzenja}<br/>
                            Broj zakasnelih vraćanja: {izvestaj.Broj_zakasnela_vracanja}<br/>
                        </Typography>
                        <Typography variant="h6" gutterBottom component="div">
                            <b>Najpopularnije knjige za odabrani period</b>
                        </Typography>
                        <Stack spacing={2} direction="row">
                            {izvestaj.Knjiga1!=0 && knjiga1!=null && <KnjigaKartica knjiga={knjiga1}/>}
                            {izvestaj.Knjiga2!=0 && knjiga2!=null &&<KnjigaKartica knjiga={knjiga2}/>}
                            {izvestaj.Knjiga3!=0 && knjiga3!=null &&<KnjigaKartica knjiga={knjiga3}/>}
                        </Stack>

                        <Typography variant="h6" gutterBottom component="div">
                            <b>Korisnici</b>
                        </Typography>
                        <Typography  gutterBottom component="div">
                            Broj korisnika: {izvestaj.Broj_korisnika}<br/>
                            Broj sumnjivih: {izvestaj.Broj_sumnjivih}<br/>
                            Broj blokiranih: {izvestaj.Broj_blokiranih}<br/>
                        </Typography>
                    </CardContent>
                </Card>
            </Grid>
            }
        </Grid>
    )
}

export default IzvestajPregled;