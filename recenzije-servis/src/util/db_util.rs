use postgres::{Client, Error, NoTls};

pub fn kreiraj_popuni_db() -> Result<(), Error> {
  let mut client = Client::connect(
      "postgresql://postgres:admin@localhost:5432/prunus-recenzije-servis-db",
      NoTls,
  )?;

  client.batch_execute("DROP TABLE IF EXISTS recenzije")?;

  client.batch_execute(
      "

      CREATE TABLE IF NOT EXISTS recenzije (
          id              SERIAL PRIMARY KEY,
          komentar VARCHAR NOT NULL,
          ocena INTEGER,
          korisnik_id          INTEGER,
          knjiga_id          INTEGER,
          obrisano        BOOLEAN,
          status          INTEGER
          )
  ",
  )?;

  client.execute(
      "INSERT INTO recenzije (komentar, ocena, korisnik_id, knjiga_id, obrisano, status) VALUES ($1, $2, $3, $4, $5, $6)",
      &[&"Sjajna knjiga.",&5.to_owned(),&1.to_owned(), &1.to_owned(), &false, &1.to_owned()],
  )?;

  client.close()?;

  Ok(())
}