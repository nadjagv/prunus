import logo from './logo.svg';
import './App.css';
import KnjigePregled from './komponente/knjige/KnjigePregled'
import Knjiga from './komponente/knjige/Knjiga'
import {Route, Routes, Link} from "react-router-dom"
import Header from "./komponente/header/Header"
import LogIn from './komponente/korisnici/LogIn';
import { useState } from 'react';
import KnjigeTabela from './komponente/knjige/KnjigeTabela';
import KorisniciTabela from './komponente/korisnici/KorisniciTabela';
import KorisnikNalogForma from './komponente/korisnici/KorisnikNalogForma';
import RezervacijeTabela from './komponente/rezervacije/RezervacijeTabela';
import PretplateTabela from './komponente/pretplate/pretplateTabela';


function App() {
  const [ulogovan, setUlogovan] = useState(false);

  function handleUlogovan(v) {
    setUlogovan(v);
  }
  return (
    <>

    <Header ulogovan = {ulogovan} handleUlogovan={handleUlogovan}/>
    <div>
      <Routes>
        <Route path="/" element = {<KnjigePregled/>} />
        <Route path="/knjige">
          <Route index element= {<KnjigePregled/>}/>
          <Route path=":id" element= {<Knjiga/>}/>
        </Route>
        <Route path="/login" element = {<LogIn handleUlogovan={handleUlogovan}/>} />
        <Route path="/registracija" element = {<KorisnikNalogForma dodavanjeMod={true}/>} />
        <Route path="/nalog" element = {<KorisnikNalogForma dodavanjeMod={false}/>} />
        <Route path="/lozinka" element = {<KorisnikNalogForma dodavanjeMod={false}/>} />

        <Route path="/uredi-knjige" element = {<KnjigeTabela/>} />
        <Route path="/korisnici" element = {<KorisniciTabela/>} />
        <Route path="/rezervacije" element = {<RezervacijeTabela/>} />
        <Route path="/pretplate" element = {<PretplateTabela/>} />

        <Route path = "*" element = {<h2>Stranica nije pronađena.</h2>}/>
      </Routes>
    </div>
    </>
  );
}

export default App;
