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
import { Button, Grid, Stack } from '@mui/material';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import useSortableData from '../../util/SortUtil';
import { format } from 'date-fns-tz';
import AuthServis from '../../servisi/AuthServis';


function Row({row, ponovoPreuzmi}) {
    function otkazi(){
        axios
          .put(`${Putanje.rezervacijeGWURL}/otkazi/${row.Id}`)
          .then((response) => {
            console.log(response.data);
            ponovoPreuzmi()
            alert("Otkazivanje uspešno!");
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }

    var istekFormatiran = format(new Date(row.DatumVremeIsteka), 'MMMM dd, yyyy')

  return (
    <React.Fragment>
        

      <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
        
        
        <TableCell>{row.Id}</TableCell>
        <TableCell>{row.KnjigaNaziv}</TableCell>
        <TableCell>{istekFormatiran}</TableCell>
        
        
        <TableCell>
          <Button
            color="primary" 
            variant="contained"
            onClick={() => otkazi()}
          >
            Otkaži
          </Button>
        </TableCell>
        
      </TableRow>
    </React.Fragment>
  );
}



export default function RezervacijeTabela() {
    const [rezervacije, setRezervacije] = useState([])

    const { items, requestSort, sortConfig } = useSortableData(rezervacije);
    const getClassNamesFor = (name) => {
        if (!sortConfig) {
        return;
        }
        return sortConfig.key === name ? sortConfig.direction : undefined;
    };

    

    useEffect(()=>{
        preuzmiSve()
        
    }, [])

    const preuzmiSve = async () => {
        axios
          .get(Putanje.rezervacijeGWURL+ "/korisnik/" + AuthServis.preuzmiKorisnika().Id)
          .then((response) => {
            console.log(response.data);
            setRezervacije(response.data)
            
          })
          .catch((error) => {
            alert("Neuspešno dobavljanje rezervacija.");
          });
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
                        onClick={() => requestSort('DatumVremeIsteka')}
                        className={getClassNamesFor('DatumVremeIsteka')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Datum isteka</b>
                </TableCell>
                
                <TableCell align="center" colSpan={3}>
                    <b>Uredi</b>
                </TableCell>
                
            </TableRow>
            </TableHead>
            {items!= null &&
            <TableBody>
            {items.map((row) => (
                <Row key={row.Id} row={row} ponovoPreuzmi={preuzmiSve}/>
            ))}

            </TableBody>
            }
        </Table>
        </TableContainer>
    </Grid>
  );
}
