#![feature(proc_macro_hygiene, decl_macro)]



#[macro_use] extern crate rocket;

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
            kontroler::kontroler::preuzmi_odobrene_po_knjizi,
            kontroler::kontroler::preuzmi_za_pregled,
            kontroler::kontroler::kreiraj,
            kontroler::kontroler::obrisi,
            kontroler::kontroler::odobri,
            kontroler::kontroler::odbij,

        ])
        .launch();
}
