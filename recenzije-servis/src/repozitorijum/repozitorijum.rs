use postgres::Row;
use postgres::{Client, NoTls};

use crate::model::recenzija::Recenzija;
use crate::model::recenzija::StatusRecenzije;

pub struct Repozitorijum {
    client: Client 
  }


impl Repozitorijum {
    pub fn new() -> Repozitorijum {
        Repozitorijum {
        client: Client::connect("postgresql://postgres:admin@localhost:5432/prunus-recenzije-servis-db",
        NoTls,).unwrap()
        }
    }

    pub fn preuzmi_sve(&mut self) -> Option<Vec<Recenzija>> {
 
        let mut recenzije: Vec<Recenzija> = vec![];
    
        let redovi = self.client.query("SELECT id, komentar, ocena, korisnik_id, knjiga_id, obrisano, status FROM recenzije WHERE obrisano = false", &[]);
    
        if redovi.is_err() {
          return None;
        }
        let unwraped_redovi = redovi.unwrap();
    
        for r in unwraped_redovi {
    
          recenzije.push(Self::mapirajNaRecenziju(r))
        }
    
        Some(recenzije)
      }

    pub fn preuzmi_po_id(&mut self, id: i32) -> Option<Recenzija> {

        let mut recenzije: Vec<Recenzija> = vec![];

        let redovi = self.client.query("SELECT id, komentar, ocena, korisnik_id, knjiga_id, obrisano, status FROM recenzije WHERE obrisano = false AND id=$1", &[&id]);

        if redovi.is_err() {
            return None;
        }
        let unwraped_redovi = redovi.unwrap();

        let r = unwraped_redovi.get(0).unwrap();

        let id: i32 = r.get(0);
        let komentar: String = r.get(1);
        let ocena: i32 = r.get(2);
        let korisnik_id: i32 = r.get(3);
        let knjiga_id: i32 = r.get(4);
        let obrisano: bool = r.get(5);
        let status: i32 = r.get(6);


        Some(

        Recenzija{
            id: (id),
            komentar: (komentar),
            ocena: (ocena),
            korisnik_id: (korisnik_id),
            knjiga_id: (knjiga_id),
            obrisano: (obrisano),
            status: (StatusRecenzije::from_i32(status))
        })
    }

    pub fn preuzmi_sve_po_knjizi(&mut self, id: i32) -> Option<Vec<Recenzija>> {
 
        let mut recenzije: Vec<Recenzija> = vec![];
    
        let redovi = self.client.query("SELECT id, komentar, ocena, korisnik_id, knjiga_id, obrisano, status FROM recenzije WHERE obrisano = false AND knjiga_id=$1", &[&id]);
    
        if redovi.is_err() {
          return None;
        }
        let unwraped_redovi = redovi.unwrap();
    
        for r in unwraped_redovi {
    
          recenzije.push(Self::mapirajNaRecenziju(r))
        }
    
        Some(recenzije)
    }

    pub fn kreiraj(&mut self, recenzija: Recenzija) -> Option<bool>{
        let redovi = self.client.query("SELECT id, komentar, ocena, korisnik_id, knjiga_id, obrisano, status FROM recenzije WHERE korisnik_id = $1 AND knjiga_id=$2", &[&recenzija.korisnik_id, &recenzija.knjiga_id]);
    
        if redovi.is_err() {
            eprintln!("{}", redovi.err().unwrap());
          return Some(false);
        }

        if !redovi.unwrap().is_empty() {
            println!("tu");
            return Some(false);
        }

        let promenjeni = self.client.execute(
            "INSERT INTO recenzije (komentar, ocena, korisnik_id, knjiga_id, obrisano, status) VALUES ($1, $2, $3, $4, $5, $6)", 
            &[&recenzija.komentar, &recenzija.ocena, &recenzija.korisnik_id, &recenzija.knjiga_id, &false, &0.to_owned()]).unwrap();
          
        if promenjeni == 0 {
            println!("tamo");
            return Some(false);
        }


        return Some(true)



    }


    pub fn obrisi(&mut self, id: i32) {
        let res = self.client.execute(
          "UPDATE recenzije SET obrisano=true WHERE id = $1",
          &[&id],);

          if res.is_err() {
            eprintln!("{}", res.err().unwrap());
        }
      }

      fn mapirajNaRecenziju(r :Row)->Recenzija{
        let id: i32 = r.get(0);
        let komentar: String = r.get(1);
        let ocena: i32 = r.get(2);
        let korisnik_id: i32 = r.get(3);
        let knjiga_id: i32 = r.get(4);
        let obrisano: bool = r.get(5);
        let status: i32 = r.get(6);

        Recenzija{
            id: (id),
            komentar: (komentar),
            ocena: (ocena),
            korisnik_id: (korisnik_id),
            knjiga_id: (knjiga_id),
            obrisano: (obrisano),
            status: (StatusRecenzije::from_i32(status))
        }
    }


}