const preuzmiKorisnika = () => {
    const korisnik = JSON.parse(localStorage.getItem("korisnik"));
    return korisnik;
  };
  
  const jeAuth = ()=>{
    const korisnik = JSON.parse(sessionStorage.getItem("korisnik"));
    return korisnik===null;
  }
  
  const postaviKorisnika = (korisnik) => {
    localStorage.setItem("korisnik", JSON.stringify(korisnik));

  };
  
  const ukloniKorisnika = () => {
    localStorage.removeItem("korisnik");
  };
  export default { preuzmiKorisnika,jeAuth, postaviKorisnika, ukloniKorisnika };