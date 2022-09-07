const preuzmiKorisnika = () => {
    const korisnik = JSON.parse(sessionStorage.getItem("korisnik"));
    return korisnik;
  };
  
  const jeAuth = ()=>{
    const korisnik = JSON.parse(sessionStorage.getItem("korisnik"));
    return korisnik===null;
  }
  
  const postaviKorisnika = (korisnik) => {
    sessionStorage.setItem("korisnik", JSON.stringify(korisnik));

  };
  
  const ukloniKorisnika = () => {
    sessionStorage.removeItem("korisnik");
  };
  export default { preuzmiKorisnika,jeAuth, postaviKorisnika, ukloniKorisnika };