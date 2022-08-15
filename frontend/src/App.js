import logo from './logo.svg';
import './App.css';
import KnjigePregled from './komponente/knjige/KnjigePregled'
import Knjiga from './komponente/knjige/Knjiga'
import {Route, Routes, Link} from "react-router-dom"
import Header from "./komponente/header/Header"
import LogIn from './komponente/korisnici/LogIn';


function App() {
  return (
    <>

    <Header/>
    <div>
      <Routes>
        <Route path="/" element = {<KnjigePregled/>} />
        <Route path="/knjige">
          <Route index element= {<KnjigePregled/>}/>
          <Route path=":id" element= {<Knjiga/>}/>
        </Route>
        <Route path="/login" element = {<LogIn/>} />

        <Route path = "*" element = {<h2>Stranica nije pronađena.</h2>}/>
      </Routes>
    </div>
    </>
  );
}

export default App;
