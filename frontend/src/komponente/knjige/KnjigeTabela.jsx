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
import { Button, Card, Grid } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';


function Row(props) {
  const { row } = props;
  const [open, setOpen] = React.useState(false);

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
        
        <TableCell>{row.Naziv}</TableCell>
        <TableCell align="right">{row.Isbn}</TableCell>
        <TableCell align="right">{row.ImeAutora} {row.PrezimeAutora}</TableCell>
        <TableCell align="right">{Object.keys(Zanr).find(key => Zanr[key] === row.Zanr)}</TableCell>
        <TableCell/>
        <TableCell>
          <IconButton
            aria-label="expand row"
            size="small"
            color = "primary"
            onClick={() => alert("promeni")}
          >
           <EditIcon></EditIcon>
          </IconButton>
        </TableCell>
        <TableCell>
          <IconButton
            aria-label="expand row"
            size="small"
            onClick={() => alert("izbrisi")}
          >
            <DeleteIcon></DeleteIcon>
          </IconButton>
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

    const navigate = useNavigate()

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
        <Button variant="contained" startIcon={<AddIcon />}
            size="large"
            color = "success"
            sx={{margin: 5}}
            onClick={() => alert("dodaj")}>
            Dodaj
        </Button>
        <TableContainer component={Paper} sx={{margin: 10, width: 0.8}}>
        <Table aria-label="collapsible table">
            <TableHead>
            <TableRow>
                <TableCell />
                <TableCell><b>Naziv</b></TableCell>
                <TableCell align="right"><b>Isbn</b></TableCell>
                <TableCell align="right"><b>Autor</b></TableCell>
                <TableCell align="right"><b>Žanr</b></TableCell>
                <TableCell align="center" colSpan={3}>
                <b>Uredi</b>
              </TableCell>
            </TableRow>
            </TableHead>
            <TableBody>
            {knjige.map((row) => (
                <Row key={row.Id} row={row} />
            ))}
            </TableBody>
        </Table>
        </TableContainer>
    </Grid>
  );
}
