import * as React from 'react';
import Box from '@mui/material/Box';
import Collapse from '@mui/material/Collapse';
import IconButton from '@mui/material/IconButton';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import { useState } from 'react';
import { useEffect } from 'react';
import Putanje from '../../konstante/Putanje';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import axios from "axios";
import { Button, Grid } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import useSortableData from '../../util/SortUtil';
import TipKorisnika from '../../enumeracije/TipKorisnika';
import KnjigaAddEditDijalog from '../knjige/KnjigaAddEditDijalog';
import { format } from 'date-fns-tz';
import AuthServis from '../../servisi/AuthServis';
import KorisnikAddEditDijalog from './KorisnikAddEditDijalog';


function Row({row, ponovoPreuzmi, admin}) {
  const [open, setOpen] = React.useState(false);

  const [dijalogOtvoren, setDijalogOtvoren] = useState(false);

    function toggleDijalogEdit(promenjeno){
        setDijalogOtvoren(!dijalogOtvoren)
        if (promenjeno){
            ponovoPreuzmi()
        }
    }

    function obrisi(){
        axios
          .delete(`${Putanje.korisniciGWURL}/${row.Id}`)
          .then((response) => {
            console.log(response.data);
            ponovoPreuzmi()
            alert("Brisanje uspešno!");
          })
          .catch((error) => {
            alert("Nije uspešno. Pokušajte ponovo.");
          });
    }

    let istekClanarineFormatiran=""
    if (row.Tip == 0){
        istekClanarineFormatiran = format(new Date(row.IstekClanarine), 'MMMM dd, yyyy')
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
        <TableCell>{row.Email}</TableCell>
        <TableCell align="right">{row.Ime}</TableCell>
        <TableCell align="right">{row.Prezime}</TableCell>
        <TableCell align="right">{Object.keys(TipKorisnika).find(key => TipKorisnika[key] === row.Tip)}</TableCell>
        <TableCell/>
        {admin &&
        <TableCell>
        <KorisnikAddEditDijalog
               otvoren={dijalogOtvoren}
               zatvoriDijalog={toggleDijalogEdit}
                admin={admin}
               korisnik = {row}
               />
          <IconButton
            aria-label="expand row"
            size="small"
            color = "primary"
            onClick={()=>toggleDijalogEdit()}
          >
           <EditIcon></EditIcon>
          </IconButton>
        </TableCell>
        }
        {admin &&
        <TableCell>
          <IconButton
            aria-label="expand row"
            size="small"
            onClick={() => obrisi()}
          >
            <DeleteIcon></DeleteIcon>
          </IconButton>
        </TableCell>
        }
      </TableRow>
      <TableRow>
        <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
          <Collapse in={open} timeout="auto" unmountOnExit>
            { row.Tip == 0 &&
            <Box sx={{ margin: 1 }}>
              <Typography variant="h6" gutterBottom component="div">
                Detalji
              </Typography>
              <Typography  gutterBottom component="div">
                Istek članarine: {istekClanarineFormatiran}
              </Typography>
            </Box>
            }
          </Collapse>
        </TableCell>
      </TableRow>
    </React.Fragment>
  );
}



export default function KorisniciTabela() {
    const [korisnici, setKorisnici] = useState([])
    const [admin, setAdmin] = useState(false)
    const [dijalogOtvoren, setDijalogOtvoren] = useState(false);

    const { items, requestSort, sortConfig } = useSortableData(korisnici);
    const getClassNamesFor = (name) => {
        if (!sortConfig) {
        return;
        }
        return sortConfig.key === name ? sortConfig.direction : undefined;
    };

    function toggleDijalog(promenjeno){
        setDijalogOtvoren(!dijalogOtvoren)
        if (promenjeno){
            preuzmiSve()
        }
    }

    useEffect(()=>{
        preuzmiSve()

        setAdmin(AuthServis.preuzmiKorisnika().Tip == 2)
        
    }, [])

    const preuzmiSve = async () => {
        axios
          .get(`${Putanje.korisniciGWURL}`)
          .then((response) => {
            console.log(response.data);
            setKorisnici(response.data)
            
          })
          .catch((error) => {
            alert("Neuspešno dobavljanje korisnika.");
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
            <KorisnikAddEditDijalog
               otvoren={dijalogOtvoren}
               zatvoriDijalog={toggleDijalog}
               dodavanjeMod={true}
               admin={admin}
               />
            {admin &&
            <Button variant="contained" startIcon={<AddIcon />}
                size="large"
                color = "success"
                sx={{margin: 5}}
                onClick={()=>toggleDijalog()}>
                Dodaj
            </Button>
            }
    
    <TableContainer component={Paper} sx={{margin: 10, width: 0.8}}>
        <Table aria-label="collapsible table">
            <TableHead>
            <TableRow>
                <TableCell />
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
                    onClick={() => requestSort('Email')}
                    className={getClassNamesFor('Email')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Email</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('Ime')}
                        className={getClassNamesFor('Ime')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Ime</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('Prezime')}
                        className={getClassNamesFor('Prezime')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Prezime</b>
                </TableCell>
                <TableCell align="right">
                <IconButton
                    aria-label="expand row"
                    size="small"
                    onClick={() => requestSort('Tip')}
                    className={getClassNamesFor('Tip')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Tip</b>
                </TableCell>
                {admin &&
                <TableCell align="center" colSpan={3}>
                <b>Uredi</b>
              </TableCell>
                }
            </TableRow>
            </TableHead>
            {items!= null &&
            <TableBody>
            {items.map((row) => (
                <Row key={row.Id} row={row} ponovoPreuzmi={preuzmiSve} admin={admin}/>
            ))}

            </TableBody>
            }
        </Table>
        </TableContainer>
    </Grid>
  );
}
