# Prunus
Veb aplikacija za biblioteku bazirana na mikroservisnoj arhitekturi.
## Funkcionalnosti

#### Neregistrovani korisnik
   - registracija
   - pregled, pretraga, sortiranje, filtriranje knjiga
   - uvid u osnovne informacije o knjigama
 
#### Registrovani korisnik
   - prijava na sistem
   - pregled i izmena profila
   - pregled, pretraga, sortiranje, filtriranje knjiga
   - uvid u osnovne informacije o knjigama
   - rezervacija knjige
   - pregled iznajmljenih knjiga
   - produženje iznajmljivanja knjige
   - ocenjivanje i komentarisanje knjige
   - pretplata za praćenje dostupnosti knjige - korisnik dobija obaveštenje na mejl kada je knjiga slobodna
   - pregled liste preporučenih knjiga (preporuka se vrši na osnovu prethodno uzetih knjiga)

#### Bibliotekar
   - prijava na sistem
   - pregled i izmena profila
   - pregled, pretraga, sortiranje, filtriranje knjiga
   - uvid u osnovne informacije o knjigama
   - CRUD operacije nad knjigama
   - iznajmljivanje knjige
   - produženje iznajmljivanja knjige
   - vraćanje knjige u biblioteku
   - pruduženje članarine
   - odobravanje i odbijanje komentara
   - pristup izveštajima o iznajmljivanju knjiga, najpopularnijim knjigama, sumnjivim korisnicima (korisnici koji nisu vratili knjigu do traženog roka)
   - slanje opomene mejlom sumnjivim korisnicima
   - blokiranje sumnjivih korisnika uz obrazloženje

#### Administrator
   - prijava na sistem
   - pregled i izmena profila
   - pregled, pretraga, sortiranje, filtriranje knjiga
   - uvid u osnovne informacije o knjigama
   - pregled, pretraga, sortiranje, filtriranje svih korisnika
   - uvid u informacije o korisniku
   - CRUD operacije nad svim korisnicima
   - pristup izveštajima o iznajmljivanju knjiga, najpopularnijim knjigama, sumnjivim korisnicima (korisnici koji nisu vratili knjigu do traženog roka)
   - slanje opomene mejlom sumnjivim korisnicima
   - blokiranje sumnjivih korisnika uz obrazloženje

## Arhitektura sistema
   - API Gateway - Go  
   - Mikroservis za knjige - CRUD sa knjigama - Go  
   - Mikroservis za korisnike - CRUD sa korisnicima, blokiranje korisnika, autorizacija - Go  
   - Mikroservis za rezervacije i iznajmljivanje - servis za rezervisanje i iznajmljivanje knjiga, pretplata na praćenje dostupnosti knjige - Go  
   - Mikroservis za komentare i ocene - servis za rad sa komentarima i ocenama - Rust  
   - Mikroservis za izveštaje - mikroservis za kreiranje i pregled izveštaja - Rust  
   - Mikroservis za slanje mejlova - mikroservis za slanje mejlova korisnicima - Go  
   - Klijentska veb aplikacija - React.js  


   - Baza podataka - PostgreSQL

*Moguće izmene servisa pre početka implementacije.

