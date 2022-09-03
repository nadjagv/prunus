import { Tabs, Tab, IconButton, Button} from "@mui/material";
import { Box } from "@mui/system";
import {React, SyntheticEvent, useEffect, useState} from "react";
import {useNavigate} from "react-router-dom"
import PersonOutlineIcon from '@mui/icons-material/PersonOutline';
import AuthServis from "../../servisi/AuthServis";

const Header = ({ulogovan, handleUlogovan}) => {
    const [value, setValue] = useState(0)
    const [korisnik, setKorisnik] = useState(null)

    useEffect(() => {
      setKorisnik(AuthServis.preuzmiKorisnika())
    }, [ulogovan])

    const navigate = useNavigate()

    const handleChange = (event, newValue) => {
      setValue(newValue);
    };

    const izlogujKorisnika = ()=>{
      AuthServis.ukloniKorisnika()
      handleUlogovan(false)
      navigate("/")
    }
  
    return (
      <Box className = "navBar">
        { korisnik==null ?
        <div width="70%">
            <Tabs value={value} onChange={handleChange} centered textColor="secondary" indicatorColor="secondary" >
              <Tab label="Sve knjige" onClick={() => navigate("/knjige")}/>
              <Tab label="Registracija" onClick={() => navigate("/registracija")}/>
              <Tab label="LogIn" onClick={() => navigate("/login")}/>
              
            </Tabs>
        </div> :
        <div width="70%">
            <Tabs value={value} onChange={handleChange} centered textColor="secondary" indicatorColor="secondary" >
              <Tab label="Sve knjige" onClick={() => navigate("/knjige")}/>
              
              { korisnik.Tip==0 && <Tab label="Rezervacija"/>}
              { korisnik.Tip==0 && <Tab label="Iznajmljivanje"/>}
              { korisnik.Tip==0 && <Tab label="Recenzije"/>}
              { korisnik.Tip==0 && <Tab label="Pretplate"/>}
              { korisnik.Tip==0 && <Tab label="Preporuka"/>}

              { korisnik.Tip==1 && <Tab label="Uredi Knjige" onClick={() => navigate("/uredi-knjige")}/>}
              { korisnik.Tip==1 && <Tab label="Iznajmljivanje"/>}
              { korisnik.Tip==1 && <Tab label="Korisnici"onClick={() => navigate("/korisnici")}/>}
              { korisnik.Tip==1 && <Tab label="Recenzije"/>}
              { korisnik.Tip==1 && <Tab label="Izveštaji"/>}

              { korisnik.Tip==2 && <Tab label="Korisnici" onClick={() => navigate("/korisnici")}/>}
              { korisnik.Tip==2 && <Tab label="Izveštaji"/>}
              
            </Tabs>
        </div>
        }
         { korisnik != null &&
          <div margin-top = "10px">
              <IconButton onClick={() => navigate ("/nalog")}> <PersonOutlineIcon/></IconButton>
              <Button color="secondary" onClick={() => izlogujKorisnika() }>Logout</Button>
          </div>
        }
        
      </Box>
    );
}

export default Header