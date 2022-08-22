use rocket::{response::content, request::{FromRequest, self}, Request, Outcome, http::{Status, RawStr}};
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{servis::servis, model::recenzija::{self, Recenzija}};

#[get("/", format = "application/json")]
pub fn preuzmi_sve() -> content::Json<String> {
  
  let mut s = servis::RecenzijaServis::new();

  let res = s.preuzmi_sve().expect("nema recenzija");

  return content::Json(Json(json!(res)).to_string());
}

#[get("/<id>", format = "application/json")]
pub fn preuzmi_po_id(id: i32) -> content::Json<String> {
  
  let mut s = servis::RecenzijaServis::new();

  let res = s.preuzmi_po_id(id).expect("nije pronadjeno");

  return content::Json(Json(json!(res)).to_string());
}

#[get("/knjiga/<id>", format = "application/json")]
pub fn preuzmi_sve_po_knjizi(id: i32) -> content::Json<String> {
  
  let mut s = servis::RecenzijaServis::new();

  let res = s.preuzmi_sve_po_knjizi(id).expect("nije pronadjeno");

  return content::Json(Json(json!(res)).to_string());
}

#[post("/", data = "<recenzija>")]
pub fn kreiraj(recenzija: Json<Recenzija>) -> content::Json<String>{
    let mut s = servis::RecenzijaServis::new();
    let res = s.kreiraj(recenzija.into_inner()).expect("neuspesno kreiranje");;

    return content::Json(Json(json!(res)).to_string());

}

#[delete("/<id>")]
pub fn obrisi(id: i32) -> content::Json<String>{
    let mut s = servis::RecenzijaServis::new();

    s.obrisi(id);

    return content::Json(String::from("Uspeh"));

}