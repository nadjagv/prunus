#![feature(proc_macro_hygiene, decl_macro)]
#[macro_use] extern crate rocket;
use rocket::config::{Config, Environment};
mod servis;
mod model;
mod kontroler;
fn main() {

    let config = Config::build(Environment::Staging)
    .address("localhost")
    .port(8001)
    .finalize().unwrap();

    rocket::custom(config)
    .mount("/", routes![kontroler::kontroler::preuzmi_izvestaj])
    .launch();
}
