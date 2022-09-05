import * as React from 'react';
import PropTypes from 'prop-types';
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
import Zanr from '../../enumeracije/Zanr';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';
import Putanje from '../../konstante/Putanje';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import KnjigaAddEditDijalog from './KnjigaAddEditDijalog';
import axios from "axios";
import { Button, Grid } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import useSortableData from '../../util/SortUtil';
import IznajmljivanjeAddDijalog from '../iznajmljivanja/IznajmljivanjeAddDijalog';


function Row({row, ponovoPreuzmi}) {
  const [open, setOpen] = React.useState(false);

  const [dijalogOtvoren, setDijalogOtvoren] = useState(false);
  const [iznajmljivanjeOtvoren, setIznajmljivanjeOtvoren] = useState(false);

    function toggleDijalogEdit(promenjeno){
        setDijalogOtvoren(!dijalogOtvoren)
        if (promenjeno){
            ponovoPreuzmi()
        }
    }

    function toggleDijalogIznajmi(promenjeno){
        setIznajmljivanjeOtvoren(!iznajmljivanjeOtvoren)
        if (promenjeno){
            ponovoPreuzmi()
        }
    }

    function obrisi(){
        axios
          .delete(`${Putanje.knjigeGWURL}/${row.Id}`)
          .then((response) => {
            console.log(response.data);
            ponovoPreuzmi()
            alert("Brisanje uspešno!");
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
        <TableCell>{row.Naziv}</TableCell>
        <TableCell align="right">{row.Isbn}</TableCell>
        <TableCell align="right">{row.ImeAutora} {row.PrezimeAutora}</TableCell>
        <TableCell align="right">{Object.keys(Zanr).find(key => Zanr[key] === row.Zanr)}</TableCell>
        <TableCell/>
        <TableCell>
        <KnjigaAddEditDijalog
               otvoren={dijalogOtvoren}
               zatvoriDijalog={toggleDijalogEdit}

               knjiga = {row}
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
        <TableCell>
          <IconButton
            aria-label="expand row"
            size="small"
            onClick={() => obrisi()}
          >
            <DeleteIcon></DeleteIcon>
          </IconButton>
        </TableCell>
        <TableCell/>
        <TableCell>
            <IznajmljivanjeAddDijalog
                otvoren={iznajmljivanjeOtvoren}
                zatvoriDijalog={toggleDijalogIznajmi}

                knjigaId = {row.Id}
                />
          <Button
            disabled= {row.TrenutnoDostupno<=0}
            color="primary" 
            variant="contained"
            onClick={() => toggleDijalogIznajmi()}
          >
            Iznajmi
          </Button>
        </TableCell>
      </TableRow>
      <TableRow>
        <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
          <Collapse in={open} timeout="auto" unmountOnExit>
            <Box sx={{ margin: 1 }}>
              <Typography variant="h6" gutterBottom component="div">
                Detalji
              </Typography>
              <Table size="small" aria-label="purchases">
                <TableHead>
                  <TableRow>
                    <TableCell>Broj strana</TableCell>
                    <TableCell>Godina nastanka</TableCell>
                    <TableCell align="right">Ukupna količina</TableCell>
                    <TableCell align="right">Trenutno dostupno</TableCell>
                    <TableCell align="right">Prosečna ocena</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow key={row.Isbn}>
                      <TableCell component="th" scope="row">
                        {row.BrojStrana}
                      </TableCell>
                      <TableCell>{row.GodinaNastanka}</TableCell>
                      <TableCell align="right">{row.UkupnaKolicina}</TableCell>
                      <TableCell align="right">{row.TrenutnoDostupno}</TableCell>
                      <TableCell align="right">{row.ProsecnaOcena}</TableCell>
                    </TableRow>
                </TableBody>
              </Table>
            </Box>
          </Collapse>
        </TableCell>
      </TableRow>
    </React.Fragment>
  );
}



export default function KnjigeTabela() {
    const [knjige, setKnjige] = useState([])
    const [dijalogOtvoren, setDijalogOtvoren] = useState(false);

    const { items, requestSort, sortConfig } = useSortableData(knjige);
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
        
    }, [])

    const preuzmiSve = async () => {
        const response = await fetch (`${Putanje.knjigeGWURL}`)
        const data = await response.json();
        console.log(data)
        setKnjige(data)
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
                    onClick={() => requestSort('Naziv')}
                    className={getClassNamesFor('Naziv')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Naziv</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('Isbn')}
                        className={getClassNamesFor('Isbn')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Isbn</b>
                </TableCell>
                <TableCell align="right">
                    <IconButton
                        aria-label="expand row"
                        size="small"
                        onClick={() => requestSort('ImeAutora')}
                        className={getClassNamesFor('ImeAutora')}
                    >
                        <SwapVertIcon></SwapVertIcon>
                    </IconButton>
                    <b>Autor</b>
                </TableCell>
                <TableCell align="right">
                <IconButton
                    aria-label="expand row"
                    size="small"
                    onClick={() => requestSort('Zanr')}
                    className={getClassNamesFor('Zanr')}
                >
                    <SwapVertIcon></SwapVertIcon>
                </IconButton>
                    <b>Žanr</b>
                </TableCell>
                <TableCell align="center" colSpan={3}>
                <b>Uredi</b>
              </TableCell>
              <TableCell align="center" colSpan={2}>
                <b>Iznajmi</b>
              </TableCell>
            </TableRow>
            </TableHead>
            {items!= null &&
            <TableBody>
            {items.map((row) => (
                <Row key={row.Id} row={row} ponovoPreuzmi={preuzmiSve} />
            ))}

            </TableBody>
            }
        </Table>
        </TableContainer>
    </Grid>
  );
}
