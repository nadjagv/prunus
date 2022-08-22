#![feature(proc_macro_hygiene, decl_macro)]

use std::env;
use dotenv::dotenv;
use rocket::{Config, config::Environment};
use rocket::http::Method;
use rocket_cors::{AllowedOrigins, CorsOptions};


#[macro_use] extern crate rocket;
extern crate base64;
extern crate dotenv;

mod util;
mod model;
mod repozitorijum;
mod servis;
mod kontroler;
fn main() {
    std::thread::spawn(|| {
        util::db_util::kreiraj_popuni_db();
    }).join().expect("Thread panicked");

    rocket::ignite()
        .mount("/",  routes![
            kontroler::kontroler::preuzmi_sve,
            kontroler::kontroler::preuzmi_po_id,
            kontroler::kontroler::preuzmi_sve_po_knjizi,
            kontroler::kontroler::kreiraj,
            kontroler::kontroler::obrisi,
        ])
        .launch();
}
