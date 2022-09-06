import React, {useEffect, useState} from "react";
import KnjigaKartica from "./KnjigaKartica";
import {useNavigate} from "react-router-dom"
import Putanje from "../../konstante/Putanje";
import { Button, IconButton, MenuItem, Select, Stack, TextField } from "@mui/material";
import SearchIcon from '@mui/icons-material/Search';

const KnjigePregled = () =>{
    const [knjige, setKnjige] = useState([])
    const [param, setParam] = useState("")
    const [pretraga, setPretraga] = useState(false)
    const [filter, setFilter] = useState(14)

    const navigate = useNavigate()

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
        <div>

                

            
            <div className = "container">
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
                </div>

                <div className="container">
                <Stack spacing={2} direction="row">
                    <Select
                        id="simple-select"
                        value={filter}
                        label="Žanr"
                        fullWidth
                        required
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
                </div>
                <div className = "container">
                    {knjige!=null ? knjige.map((knjiga) => (
                        <div onClick={() => navigate("/knjige/" + knjiga.Id)}>
                            {((filter==14) || (filter!=14 && filter==knjiga.Zanr)) &&
                            <KnjigaKartica knjiga={knjiga} key = {knjiga.Id}/>
                            }
                        </div>
                    )) : 
                        <div className = "empty">
                            <h2>Nema knjiga za pregled.</h2>
                        </div>
                    }
                </div>

        </div>
    )
}

export default KnjigePregled