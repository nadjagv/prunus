import * as React from 'react';
import IconButton from '@mui/material/IconButton';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { useState } from 'react';
import { useEffect } from 'react';
import Putanje from '../../konstante/Putanje';
import axios from "axios";
import { Button, Grid, MenuItem, Select, Stack } from '@mui/material';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import useSortableData from '../../util/SortUtil';
import { format } from 'date-fns-tz';
import AuthServis from '../../servisi/AuthServis';
import RecenzijaAddDijalog from '../recenzije/RecenzijaAddDijalog';


function Row({row, ponovoPreuzmi, clan, filter}) {
    const [dijalogOtvoren, setDijalogOtvoren] = useState(false);
    const [postojiRecenzija, setPostojiRecenzija] = useState(false);

    useEffect(()=>{
      if(clan){
        proveriRecenzijaPostoji()
      }
      
    }, [])

    function proveriRecenzijaPostoji(){
      axios
          .get(`${Putanje.recenzijeGWURL}/postoji/${row.KorisnikId}/${row.KnjigaId}`)
          .then((response) => {
            setPostojiRecenzija(response.data);
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }

    function toggleDijalog(promenjeno){
        setDijalogOtvoren(!dijalogOtvoren)
        if(promenjeno){
          proveriRecenzijaPostoji()
        }
    }

    function vrati(){
        axios
          .post(`${Putanje.iznajmljivanjaGWURL}/vrati`, row)
          .then((response) => {
            console.log(response.data);
            ponovoPreuzmi()
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }

    function produzi(){
        axios
          .put(`${Putanje.iznajmljivanjaGWURL}/produzi/${row.Id}`)
          .then((response) => {
            console.log(response.data);
            ponovoPreuzmi()
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }

    var rokDate = new Date(row.RokVracanja);

    var danas = new Date();
    var prosaoRok = false

    if(rokDate.setHours(0,0,0,0) < danas.setHours(0,0,0,0)) {
        prosaoRok=true
    }

    var iznajmljivanjeFormatiran = format(new Date(row.DatumVremeIznajmljivanja), 'MMMM dd, yyyy')
    var rokFormatiran = format(new Date(row.RokVracanja), 'MMMM dd, yyyy')
   
    var vracanjeFormatiran = format(new Date(row.DatumVremeVracanja), 'MMMM dd, yyyy')
    var vraceno = true
    if (new Date(row.DatumVremeVracanja).getFullYear() == 1){
        vraceno = false
        vracanjeFormatiran=""
    }

  return (
    <React.Fragment>
        
        {((filter==0) || (filter==1 && row.Aktivno==true) || (filter==2 && row.Aktivno==false))&&
      <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
        
        
        <TableCell>{row.Id}</TableCell>
        <TableCell>{row.KorisnikEmail}</TableCell>
        <TableCell>{row.KnjigaNaziv}</TableCell>
        <TableCell>{iznajmljivanjeFormatiran}</TableCell>
        <TableCell>{rokFormatiran}</TableCell>
        <TableCell>{vracanjeFormatiran}</TableCell>
        
        {!clan &&
        <TableCell>
          <Button
            disabled= {vraceno}
            color="primary" 
            variant="contained"
            onClick={() => vrati()}
          >
            Vrati
          </Button>
        </TableCell>
        }

        {clan &&
        <TableCell>
            <RecenzijaAddDijalog
               otvoren={dijalogOtvoren}
               zatvoriDijalog={toggleDijalog}
               knjigaId = {row.KnjigaId}
               />
          <Button
            disabled = {postojiRecenzija}
            color="primary" 
            variant="contained"
            onClick={() => toggleDijalog()}
          >
            Recenziraj
          </Button>
        </TableCell>
        }

        <TableCell>
          <Button
            disabled= {prosaoRok || row.Produzeno || vraceno}
            color="primary" 
            variant="contained"
            onClick={() => produzi()}
          >
            Produži
          </Button>
        </TableCell>
        
      </TableRow>
    }
    </React.Fragment>
  );
}



export default function IznajmljivanjaTabela({istorija}) {
    const [iznajmljivanja, setIznajmljivanja] = useState([])
    const [clan, setClan] = useState(false)
    const [filter, setFilter] = useState(0)

    const { items, requestSort, sortConfig } = useSortableData(iznajmljivanja);
    const getClassNamesFor = (name) => {
        if (!sortConfig) {
        return;
        }
        return sortConfig.key === name ? sortConfig.direction : undefined;
    };

    

    useEffect(()=>{
        preuzmiSve()
        
    }, [istorija])

    const preuzmiSve = async () => {
        var korisnik = AuthServis.preuzmiKorisnika()
        if (korisnik.Tip == 0){
            setClan(true)
            if (istorija){
              axios
              .get(Putanje.iznajmljivanjaGWURL+ "/sve-korisnik/" + korisnik.Id)
              .then((response) => {
                  console.log(response.data);
                  setIznajmljivanja(response.data)
                  
              })
              .catch((error) => {
                  alert("Neuspešno dobavljanje iznajmljivanja.");
              });

            }else{
              axios
              .get(Putanje.iznajmljivanjaGWURL+ "/aktivna-korisnik/" + korisnik.Id)
              .then((response) => {
                  console.log(response.data);
                  setIznajmljivanja(response.data)
                  
              })
              .catch((error) => {
                  alert("Neuspešno dobavljanje iznajmljivanja.");
              });
            }
        }
        else{
            axios
            .get(Putanje.iznajmljivanjaGWURL)
            .then((response) => {
                console.log(response.data);
                setIznajmljivanja(response.data)
                
            })
            .catch((error) => {
                alert("Neuspešno dobavljanje iznajmljivanja.");
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

    {(!clan || istorija) &&
    <Stack spacing={2} direction="row" sx={{margin:3}}>
        <Select
            id="simple-select"
            value={filter}
            label="Status"
            fullWidth
            onChange={(e) => {setFilter(e.target.value);}}
            >
            <MenuItem value={0}>Sve</MenuItem>
            <MenuItem value={1}>Aktivna</MenuItem>
            <MenuItem value={2}>Neaktivna</MenuItem>
            
        </Select>
        </Stack>
    }
    <TableContainer component={Paper} sx={{margin: 10, width: 0.8}}>
        <Table aria-label="collapsible table">
            <TableHead>
            <TableRow>
                
                <TableCell>
                <IconButton
                    size="small"
                    onClick={() => requestSort('Id')}
                    className={getClassNamesFor('Id')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Id</b>
                </TableCell>
                <TableCell>
                <IconButton
                    size="small"
                    onClick={() => requestSort('KorisnikEmail')}
                    className={getClassNamesFor('KorisnikEmail')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Korisnik</b>
                </TableCell>
                <TableCell>
                <IconButton
                    aria-label="expand row"
                    size="small"
                    onClick={() => requestSort('KnjigaNaziv')}
                    className={getClassNamesFor('KnjigaNaziv')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Naziv knjige</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('DatumVremeIznajmljivanja')}
                        className={getClassNamesFor('DatumVremeIznajmljivanja')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Datum iznajmljivanja</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('RokVracanja')}
                        className={getClassNamesFor('RokVracanja')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Rok za vraćanje</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('DatumVremeVracanja')}
                        className={getClassNamesFor('DatumVremeVracanja')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Datum vraćanja</b>
                </TableCell>
                
                <TableCell align="center" colSpan={3}>
                    <b>Uredi</b>
                </TableCell>
                
            </TableRow>
            </TableHead>
            {items!= null &&
            <TableBody>
            {items.map((row) => (
                <Row key={row.Id} row={row} ponovoPreuzmi={preuzmiSve} clan={clan} filter={filter}/>
            ))}

            </TableBody>
            }
        </Table>
        </TableContainer>
    </Grid>
  );
}
