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
    Kreirano=0,
    Odobreno=1,
    Odbijeno=2
}