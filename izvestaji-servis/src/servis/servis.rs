extern crate reqwest; 
use std::io::Read;
use crate::model::izvestaj::Izvestaj;
use std::collections::HashMap;

pub struct IzvestajServis{}

impl IzvestajServis{
    pub fn new() -> IzvestajServis{
        IzvestajServis{}
    }

    pub fn napravi_izvestaj(&mut self, pocetak: i64, kraj: i64)->Izvestaj{
        let url = format!("http://localhost:8083/iznajmljivanja/izmedju-datuma/sve?pocetak={}&kraj={}", pocetak, kraj);
        let mut res = reqwest::blocking::get(&url).unwrap();
        let mut body = String::new();
        res.read_to_string(&mut body);

        let parsed = json::parse(&body).unwrap();

        let broj_elem = parsed.len() as i32;

        let mut broj_iznajmljivanja= broj_elem;
        let mut broj_zakasnela_vracanja =0;
        let mut broj_produzenja=0;

        let mut knjiga1=0;
        let mut knjiga2=0;
        let mut knjiga3=0;

        let mut broj_korisnika=0;
        let mut broj_sumnjivih=0;
        let mut broj_blokiranih=0;

        if broj_iznajmljivanja != 0{
            let mut iznajmljene_knjige_ids:HashMap<i32, i32> = HashMap::new();
    
            for i in 0..broj_elem{
                let indeks = i as usize;
                if parsed[indeks]["ZakasneloVracanje"] == true{
                    broj_zakasnela_vracanja += 1;
                }
    
                if parsed[indeks]["Produzeno"] == true{
                    broj_produzenja += 1;
                }
                let knjiga_id = parsed[indeks]["KnjigaId"].as_i32().unwrap();
    
                *iznajmljene_knjige_ids.entry(knjiga_id).or_insert(0) += 1;
    
            }
    
            let mut knjige_vec: Vec<_> = iznajmljene_knjige_ids.iter().collect();
            knjige_vec.sort_by(|a, b| b.1.cmp(a.1));
    
            knjiga1 = *knjige_vec[0].0;
            if knjige_vec.len() >= 2{
                knjiga2 = *knjige_vec[1].0;
            }

            if knjige_vec.len() >= 3{
                knjiga3 = *knjige_vec[2].0;
            }
  
        }

        //deo za korisnike
        let url = format!("http://localhost:8082/");
        let mut res = reqwest::blocking::get(&url).unwrap();
        let mut body = String::new();
        res.read_to_string(&mut body);

        let parsed = json::parse(&body).unwrap();
        let broj_elem = parsed.len() as i32;
        broj_korisnika=broj_elem;
        for i in 0..broj_elem{
            let indeks = i as usize;
                if parsed[indeks]["Sumnjiv"] == true{
                    broj_sumnjivih += 1;
                }
    
                if parsed[indeks]["Blokiran"] == true{
                    broj_blokiranih += 1;
                }
        }
        return Izvestaj {
                 broj_iznajmljivanja: broj_iznajmljivanja,
                 broj_zakasnela_vracanja: broj_zakasnela_vracanja,
                 broj_produzenja: broj_produzenja,

                 knjiga1: knjiga1,
                 knjiga2: knjiga2,
                 knjiga3: knjiga3,

                 broj_korisnika: broj_korisnika,
                 broj_sumnjivih: broj_sumnjivih,
                 broj_blokiranih: broj_blokiranih,
            }
        

    }

}