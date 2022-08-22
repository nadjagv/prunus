use crate::model::recenzija;
use crate::model::recenzija::Recenzija;
use crate::model::recenzija::StatusRecenzije;

use crate::repozitorijum;

pub struct RecenzijaServis{
    repo: repozitorijum::repozitorijum::Repozitorijum
}

impl RecenzijaServis{
    pub fn new() -> RecenzijaServis{
        RecenzijaServis{repo: repozitorijum::repozitorijum::Repozitorijum::new()}
    }

    pub fn preuzmi_sve(&mut self) -> Result<Vec<Recenzija>, String> {
        match self.repo.preuzmi_sve() {
          Some(recenzije) => Ok(recenzije),
          None => Err("Nema recenzija.".to_string())
        }
    }

    pub fn preuzmi_po_id(&mut self, id: i32) -> Result<Recenzija, String> {
        match self.repo.preuzmi_po_id(id) {
            Some(recenzija) => Ok(recenzija),
            None => Err(format!("Nije pronađena recenzija sa id {}", id))
        }
    }

    pub fn preuzmi_sve_po_knjizi(&mut self, id: i32) -> Result<Vec<Recenzija>, String> {
        match self.repo.preuzmi_sve_po_knjizi(id) {
          Some(recenzije) => Ok(recenzije),
          None => Err("Nema recenzija.".to_string())
        }
    }

    pub fn kreiraj(&mut self, recenzija: Recenzija) -> Result<String, String> {
        match self.repo.kreiraj(recenzija) {
          Some(true) => Ok("Uspeh".to_string()),
          Some(false) => Err("Neuspešno kreiranje.".to_string()),
          None => Err("Neuspešno kreiranje.".to_string())
        }
      }
    
      pub fn obrisi(&mut self, id: i32) {
        self.repo.obrisi(id);
      }

      pub fn promeni_status(&mut self, id: i32, status: i32) {
        self.repo.promeni_status(id, status);
      }
}