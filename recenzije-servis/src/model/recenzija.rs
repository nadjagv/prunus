use serde_derive::{Serialize, Deserialize};

#[derive(Serialize, Clone, Deserialize)]
pub struct Recenzija {
  pub id: i32,
  pub komentar: String,
  pub ocena: i32,
  pub korisnik_id: i32,
  pub knjiga_id: i32,
  pub obrisano: bool,
  pub status: StatusRecenzije
}

#[derive(Serialize, Clone, Deserialize)]
pub enum StatusRecenzije{
    KREIRANO=0,
    ODOBRENO=1,
    ODBIJENO=2
}

impl StatusRecenzije {
    pub fn from_i32(status: i32) -> StatusRecenzije {
        match status {
            0 => StatusRecenzije::KREIRANO,
            1 => StatusRecenzije::ODOBRENO,
            2 => StatusRecenzije::ODBIJENO,
            _ => panic!("Nepoznata vrednost: {}", status),
        }
    }
}