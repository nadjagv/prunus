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
import { Box, Button, Collapse, Grid, Stack, Typography } from '@mui/material';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import useSortableData from '../../util/SortUtil';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';


function Row({row, ponovoPreuzmi, admin}) {
    const [open, setOpen] = React.useState(false);
  
  
      function odobri(){
        axios
            .put(`${Putanje.recenzijeGWURL}/odobri/${row.Id}`)
            .then((response) => {
              console.log(response.data);
              ponovoPreuzmi()
              alert("Odobreno.");
            })
            .catch((error) => {
              alert("Nije uspešno. Pokušajte ponovo.");
            });
      }
  
      function odbij(){
          axios
            .put(`${Putanje.recenzijeGWURL}/odbij/${row.Id}`)
            .then((response) => {
              console.log(response.data);
              ponovoPreuzmi()
              alert("Odbijeno.");
            })
            .catch((error) => {
              alert("Nije uspešno. Pokušajte ponovo.");
            });
      }
  
      
  
    return (
      <React.Fragment>
          
  
        <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
          <TableCell>
            <IconButton
              aria-label="expand row"
              size="small"
              onClick={() => setOpen(!open)}
            >
              {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
            </IconButton>
          </TableCell>
          
          
          <TableCell>{row.Id}</TableCell>
          <TableCell>{row.KnjigaNaziv}</TableCell>
          <TableCell>{row.KorisnikEmail}</TableCell>
          <TableCell>{row.Ocena}</TableCell>
          
          <TableCell>
            <Button
                color="success" 
                variant="contained"
                onClick={() => odobri()}
            >
                Odobri
            </Button>
          </TableCell>
          <TableCell>
            <Button
                color="error" 
                variant="contained"
                onClick={() => odbij()}
            >
                Odbij
            </Button>
          </TableCell>
          
        </TableRow>
        <TableRow>
          <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
            <Collapse in={open} timeout="auto" unmountOnExit>
              
              <Box sx={{ margin: 1 }}>
              
                <Typography variant="h6" gutterBottom component="div">
                  Komentar
                </Typography>
                <Typography  gutterBottom component="div">
                  {row.Komentar}
                </Typography>
              </Box>
              
            </Collapse>
          </TableCell>
        </TableRow>
      </React.Fragment>
    );
  }



export default function RecenzijeTabela() {
    const [recenzije, setRecenzije] = useState([])

    const { items, requestSort, sortConfig } = useSortableData(recenzije);
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
            .get(Putanje.recenzijeGWURL+ "/pregled/sve/")
            .then((response) => {
                console.log(response.data);
                setRecenzije(response.data)
                
            })
            .catch((error) => {
                alert("Neuspešno dobavljanje recenzija.");
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
                <TableCell/>
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
                        onClick={() => requestSort('KorisnikEmail')}
                        className={getClassNamesFor('KorisnikEmail')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Email korisnika</b>
                </TableCell>

                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('Ocena')}
                        className={getClassNamesFor('Ocena')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Ocena</b>
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
