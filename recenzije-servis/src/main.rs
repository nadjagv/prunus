mod util;
mod model;

fn main() {
    std::thread::spawn(|| {
        util::db_util::kreiraj_popuni_db();
    }).join().expect("Thread panicked");
}
