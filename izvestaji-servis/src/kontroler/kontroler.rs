use rocket::{response::content, request::{FromRequest, self}, Outcome, http::{Status, RawStr}};
use rocket_contrib::json::Json;
use serde_json::json;
use reqwest::Client;
use std::error::Error;

use crate::{servis::servis, model::izvestaj::{self, Izvestaj}};

#[get("/?<pocetak>&<kraj>", format = "application/json")]
pub fn preuzmi_izvestaj(pocetak: i64, kraj:i64) -> content::Json<String> {
  
    let mut s = servis::IzvestajServis::new();

    let res =s.napravi_izvestaj(pocetak, kraj);

  return content::Json(Json(json!(res)).to_string());
}