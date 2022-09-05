use serde_derive::{Serialize, Deserialize};

#[derive(Serialize, Clone, Deserialize)]
pub struct Izvestaj {
  pub broj_iznajmljivanja: i32,
  pub broj_zakasnela_vracanja: i32,
  pub broj_produzenja: i32,

  pub knjiga1: i32,
  pub knjiga2: i32,
  pub knjiga3: i32,

  pub broj_korisnika: i32,
  pub broj_sumnjivih: i32,
  pub broj_blokiranih: i32,
}