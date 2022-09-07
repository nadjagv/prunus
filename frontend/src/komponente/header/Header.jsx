import { Tabs, Tab, IconButton, Button, MenuItem, Menu} from "@mui/material";
import { Box } from "@mui/system";
import {React, SyntheticEvent, useEffect, useState} from "react";
import {useNavigate} from "react-router-dom"
import PersonOutlineIcon from '@mui/icons-material/PersonOutline';
import AuthServis from "../../servisi/AuthServis";

const Header = ({ulogovan, handleUlogovan}) => {
    const [value, setValue] = useState(0)
    const [korisnik, setKorisnik] = useState(null)
    const [anchorEl, setAnchorEl] = useState(null);
    const open = Boolean(anchorEl);

    useEffect(() => {
      setKorisnik(AuthServis.preuzmiKorisnika())
    }, [ulogovan])

    const navigate = useNavigate()

    const handleChange = (event, newValue) => {
      setValue(newValue);
    };

    const handleNalogClick = (event) => {
      setAnchorEl(event.currentTarget);
  };

    const handleNalogClose = () => {
      setAnchorEl(null);
  };

    const mojNalog = () => {
      setAnchorEl(null);
      navigate("/nalog")
  };

  const promenaLozinke = () => {
    setAnchorEl(null);
    navigate("/lozinka")
  };

    const izlogujKorisnika = ()=>{
      AuthServis.ukloniKorisnika()
      setAnchorEl(null);
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
              
              { korisnik.Tip==0 && <Tab label="Rezervacije" onClick={() => navigate("/rezervacije")}/>}
              { korisnik.Tip==0 && <Tab label="Iznajmljivanja"onClick={() => navigate("/iznajmljivanja")}/>}
              { korisnik.Tip==0 && <Tab label="Dnevnik čitanja"onClick={() => navigate("/istorija")}/>}
              { korisnik.Tip==0 && <Tab label="Pretplate" onClick={() => navigate("/pretplate")}/>}
              { korisnik.Tip==0 && <Tab label="Preporuka"onClick={() => navigate("/preporuka")}/>}

              { korisnik.Tip==1 && <Tab label="Uredi Knjige" onClick={() => navigate("/uredi-knjige")}/>}
              { korisnik.Tip==1 && <Tab label="Iznajmljivanja"onClick={() => navigate("/iznajmljivanja")}/>}
              { korisnik.Tip==1 && <Tab label="Korisnici"onClick={() => navigate("/korisnici")}/>}
              { korisnik.Tip==1 && <Tab label="Recenzije" onClick={() => navigate("/recenzije")}/>}
              { korisnik.Tip==1 && <Tab label="Izveštaji" onClick={() => navigate("/izvestaji")}/>}

              { korisnik.Tip==2 && <Tab label="Korisnici" onClick={() => navigate("/korisnici")}/>}
              { korisnik.Tip==2 && <Tab label="Izveštaji" onClick={() => navigate("/izvestaji")}/>}
              
            </Tabs>
        </div>
        }
         { korisnik != null &&
          <div margin-top = "10px">
              <IconButton 
                id='basic-button'
                onClick={handleNalogClick}
                aria-controls={open ? 'basic-menu' : undefined}
                aria-haspopup="true"
                aria-expanded={open ? 'true' : undefined}> 
                <PersonOutlineIcon/>
              </IconButton>
              <Menu
              id="basic-menu"
              anchorEl={anchorEl}
              open={open}
              onClose={handleNalogClose}
              MenuListProps={{
                'aria-labelledby': 'basic-button',
              }}
            >
              <MenuItem onClick={mojNalog}>Moj nalog</MenuItem>
              <MenuItem onClick={promenaLozinke}>Promena lozinke</MenuItem>
              <MenuItem onClick={() => izlogujKorisnika() }>Logout</MenuItem>
            </Menu>
          </div>
        }
        
      </Box>
    );
}

export default Header