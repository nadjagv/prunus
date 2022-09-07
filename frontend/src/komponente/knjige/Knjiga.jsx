import { TryRounded } from "@mui/icons-material";
import { Box, Button, Card, CardContent, Grid, List, Rating } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Zanr from "../../enumeracije/Zanr"
import Putanje from "../../konstante/Putanje";
import AuthServis from "../../servisi/AuthServis";
import axios from "axios";

const Knjiga = () =>{
    const { id } = useParams()

    const [knjiga, setKnjiga] = useState({ProsecnaOcena:5})
    const [komentari, setKomentari] = useState([])
    const [korisnik, setKorisnik] = useState({})
    const [pretplataPostoji, setPretplataPostoji] = useState(false)
    const [rezervacijaPostoji, setRezervacijaPostoji] = useState(false)
    const [maksimalnoRezervisao, setMaksimalnoRezervisao] = useState(false)

    useEffect(()=>{
        setKorisnik(AuthServis.preuzmiKorisnika())     
    }, [])

    useEffect(()=>{
        preuzmiPoId()
        preuzmiOdobreneKomentare()
        if (korisnik != null){
            preuzmiPretplacen()
            preuzmiRezervaciju()
        }
        
        
        
    }, [korisnik])

    const preuzmiPoId = async () => {
        const response = await fetch (`${Putanje.knjigeGWURL}/${id}`)
        const data = await response.json();
        console.log(data)
        setKnjiga(data)
    }

    const preuzmiOdobreneKomentare = async () => {
        const response = await fetch (`${Putanje.recenzijeGWURL}/knjiga-odobreni/${id}`)
        const data = await response.json();
        console.log(data)
        setKomentari(data)
    }

    const preuzmiPretplacen = ()=>{
        
        if (korisnik.Id != undefined){
            axios
          .get(`${Putanje.knjigeGWURL}/pretplata/knjiga-korisnik/${id}/${korisnik.Id}`)
          .then((response) => {
            console.log(response.data)
            setPretplataPostoji(response.data.Id!=0)
          })
          .catch((error) => {
            console.log("Pretplata ne postoji.")
          });
        }
        
    }

    const preuzmiRezervaciju = ()=>{
        if (korisnik.Id != undefined){
            axios
          .get(`${Putanje.rezervacijeGWURL}/knjiga-korisnik/${id}/${korisnik.Id}`)
          .then((response) => {
            console.log(response.data)
            setRezervacijaPostoji(response.data.Id!=0)
            setMaksimalnoRezervisao(false)
          })
          .catch((error) => {
            setMaksimalnoRezervisao(true)
          });
        }
        
    }

    const pretplatiSe=() =>{
        let pretplataDto = {
            KorisnikId: parseInt(korisnik.Id),
            KorisnikEmail: korisnik.Email,
            KnjigaId: parseInt(id),
            KnjigaNaziv: knjiga.Naziv,
        }
        axios
          .post(`${Putanje.knjigeGWURL}/pretplata`, pretplataDto)
          .then((response) => {
            console.log(response.data)
            setPretplataPostoji(true)
          })
          .catch((error) => {
            alert("Neuspešna pretplata.")
          });
    }


    const rezervisi=() =>{
        let rezervacijaDto = {
            KorisnikId: parseInt(korisnik.Id),
            KnjigaId: parseInt(id),
        }
        axios
          .post(`${Putanje.rezervacijeGWURL}`, rezervacijaDto)
          .then((response) => {
            setRezervacijaPostoji(true)
          })
          .catch((error) => {
            alert("Neuspešna rezervacija.")
          });
    }


    return (
        <>
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
                        <Rating name="read-only" value={knjiga.ProsecnaOcena} readOnly precision={0.5} />
                        <h4>ISBN: {knjiga.Isbn}</h4>
                        <h4>Zanr: {Object.keys(Zanr).find(key => Zanr[key] === knjiga.Zanr)}</h4>

                        <br/>
                        <p>{knjiga.Opis}</p>

                        <br/>
                        <p> Broj strana: {knjiga.BrojStrana}</p>
                        <p> Godina nastanka: {knjiga.GodinaNastanka}</p>

                        <br/>

                        {korisnik!=null && korisnik.Tip==0 && 
                        <div>
                            {knjiga.TrenutnoDostupno<=0 && !pretplataPostoji &&
                            <div>
                                
                                <p>Knjiga trenutno nije dostupna. Moguće je praćenje dostupnosti klikom na dugme PRETPLATI SE.</p>
                                <br/>
                                <Button 
                                    color="primary" 
                                    variant="contained"
                                    onClick={() => pretplatiSe()}>
                                        Pretplati se
                                </Button>
                            </div>
                            }
                            {knjiga.TrenutnoDostupno<=0 && pretplataPostoji &&
                            <div>
                                
                                <p><b>Pretplaćeni ste</b> na ovu knjigu. Otkazivanje pretplate možete izvršiti na tabu PRETPLATE.</p>
                                <br/>
                            </div>
                            }
                            { knjiga.TrenutnoDostupno>0 && !rezervacijaPostoji &&
                            <div>
                                <p>Knjiga je dostupna.</p>
                                <br/>
                                {!maksimalnoRezervisao &&
                                <Button 
                                    color="primary" 
                                    variant="contained"
                                    onClick={() => rezervisi()}>
                                        Rezerviši
                                </Button>
                                }
                            </div>
                            }
                            { knjiga.TrenutnoDostupno>0 && rezervacijaPostoji &&
                            <div>
                                <p>Knjiga je dostupna.</p>
                                <br/>
                                <div>
                                <p><b>Rezervisali ste</b> ovu knjigu. Detalje pogledajte na tabu REZERVACIJA.</p>
                                <br/>
                            </div>
                            </div>
                            }
                        
                        </div>
                        }
                    </CardContent>
                </Card>
            </Grid>
          </Grid>
        </Box>

        <Grid
        container
        spacing={50}
        direction="column"
        alignItems="center"
        justify="center"
        style={{ minHeight: '100vh' }}
        >

            <Grid item xs={3} sx={{marginRight: 10, width: 1/2}}>

                <List  style={{maxHeight: '500px', overflow: 'auto', display:'block'}}>
                {komentari.length > 0 ? (
                    <>
                            {komentari.map((komentar) => (
                                <Card sx={{marginTop: 3}} key = {komentar.Id}>
                                    <CardContent>
                                        <Rating name="read-only" value={komentar.Ocena} readOnly />
                                        <p margin="10px">{komentar.Komentar}</p>
                                    </CardContent>
                                </Card>
                            ))}
                    </>
                    ):(
                            <h2>Nema komentara.</h2>
                        
                    )}
                </List>
            </Grid>      
        </Grid>
        </>
      );
}

export default Knjiga