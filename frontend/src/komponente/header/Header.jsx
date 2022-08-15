import { Tabs, Tab, IconButton, Button} from "@mui/material";
import { Box } from "@mui/system";
import {React, SyntheticEvent, useState} from "react";
import {useNavigate} from "react-router-dom"
import PersonOutlineIcon from '@mui/icons-material/PersonOutline';

const Header = () => {
    const [value, setValue] = useState(0);

    const navigate = useNavigate()

    const handleChange = (event, newValue) => {
      setValue(newValue);
    };
  
    return (
      <Box className = "navBar">
        <div width="70%">
            <Tabs value={value} onChange={handleChange} centered textColor="secondary"
    indicatorColor="secondary" >
            <Tab label="Sve knjige" onClick={() => navigate("/knjige")}/>
            <Tab label="Item Two" />
            <Tab label="Item Three" />
            </Tabs>
        </div>

        <div margin-top = "10px">
            <Tab label="Registracija" onClick={() => navigate("/registracija")}/>
        
            <Tab label="LogIn" onClick={() => navigate("/login")}/>
        
            <IconButton onClick={() => navigate ("/nalog")}> <PersonOutlineIcon/></IconButton>
        </div>
        
      </Box>
    );
}

export default Header