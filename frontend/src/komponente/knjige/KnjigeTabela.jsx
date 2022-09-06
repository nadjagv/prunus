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
import { Button, Grid, MenuItem, Select, Stack, TextField } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import SwapVertIcon from '@mui/icons-material/SwapVert';
import useSortableData from '../../util/SortUtil';
import IznajmljivanjeAddDijalog from '../iznajmljivanja/IznajmljivanjeAddDijalog';
import SearchIcon from '@mui/icons-material/Search';


function Row({row, ponovoPreuzmi, filter}) {
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
        
      {((filter==14) || (filter!=14 && filter==row.Zanr)) &&
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
      }
      {((filter==14) || (filter!=14 && filter==row.Zanr)) &&
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
    }
    </React.Fragment>
  );
}



export default function KnjigeTabela() {
    const [knjige, setKnjige] = useState([])
    const [dijalogOtvoren, setDijalogOtvoren] = useState(false);
    const [param, setParam] = useState("")
    const [pretraga, setPretraga] = useState(false)
    const [filter, setFilter] = useState(14)

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
      if(!pretraga){
          preuzmiSve()
      }
  }, [pretraga])

    const preuzmiSve = async () => {
        const response = await fetch (`${Putanje.knjigeGWURL}`)
        const data = await response.json();
        console.log(data)
        setKnjige(data)
    }

    const pretrazi = async () => {
      setPretraga(true)
      const response = await fetch (`${Putanje.knjigeGWURL}/pretrazi/${param}`)
      const data = await response.json();
      console.log(data)
      setKnjige(data)
  }

  const ponisti=()=>{
      setParam("")
      setPretraga(false)
      
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

            <Stack spacing={2} direction="row">
              <TextField
              margin="normal"
              label="Pretraga"
              placeholder="Pretraži"
              fullWidth
              value={param}
              onChange={(e) => {
                  setParam(e.target.value);
              }}
              ></TextField>

              <IconButton
                  aria-label="expand row"
                  size="large"
                  onClick={() => pretrazi()}
                  
              >
                  <SearchIcon></SearchIcon>
              </IconButton>

              <Button 
                  color="primary" 
                  variant="contained"
                  onClick={() => ponisti()}>
                      Poništi pretragu
                  </Button>
              </Stack>


              <Stack spacing={2} direction="row">
                    <Select
                        id="simple-select"
                        value={filter}
                        label="Žanr"
                        fullWidth
                        onChange={(e) => {setFilter(e.target.value);}}
                        >
                        <MenuItem value={14}>Sve</MenuItem>
                        <MenuItem value={0}>Naučna fantastika</MenuItem>
                        <MenuItem value={1}>Ljubavni</MenuItem>
                        <MenuItem value={2}>Klasik</MenuItem>
                        <MenuItem value={3}>Horor</MenuItem>
                        <MenuItem value={4}>Triler</MenuItem>
                        <MenuItem value={5}>Avantura</MenuItem>
                        <MenuItem value={6}>Biografija</MenuItem>
                        <MenuItem value={7}>Popularna psihologija</MenuItem>
                        <MenuItem value={8}>Opšta interesovanja</MenuItem>
                        <MenuItem value={9}>Stručna literatura</MenuItem>
                        <MenuItem value={10}>Strani jezik</MenuItem>
                        <MenuItem value={11}>Poezija</MenuItem>
                        <MenuItem value={12}>Dečije</MenuItem>
                        <MenuItem value={13}>Ostalo</MenuItem>
                        
                    </Select>
                    </Stack>
    
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
                <Row key={row.Id} row={row} ponovoPreuzmi={preuzmiSve} filter={filter}/>
            ))}

            </TableBody>
            }
        </Table>
        </TableContainer>
    </Grid>
  );
}
