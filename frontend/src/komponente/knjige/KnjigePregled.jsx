import React, {useEffect, useState} from "react";
import KnjigaKartica from "./KnjigaKartica";
import {useNavigate} from "react-router-dom"
import Putanje from "../../konstante/Putanje";
import { Button, IconButton, Stack, TextField } from "@mui/material";
import SearchIcon from '@mui/icons-material/Search';

const KnjigePregled = () =>{
    const [knjige, setKnjige] = useState([])
    const [param, setParam] = useState("")
    const [pretraga, setPretraga] = useState(false)

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
                <div className = "container">
                    {knjige!=null ? knjige.map((knjiga) => (
                        <div onClick={() => navigate("/knjige/" + knjiga.Id)}>
                            <KnjigaKartica knjiga={knjiga} key = {knjiga.Id}/>
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