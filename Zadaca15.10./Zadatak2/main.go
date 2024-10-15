// Napiši strukturu koja predstavlja adresu, koja sadrži polja za grad i ulicu.
// Zatim napiši strukturu koja predstavlja osobu, koja sadrži ime, godine i adresu.
// Kreiraj metodu koja ispisuje puni opis osobe, uključujući njezinu adresu.
// (Sva polja) u formatu: Ime Prezime, 20 godina, živi u Grad, Ulica.

type Adresa struct {
	grad  string
	ulica string
}

type Osoba struct {
	ime     string
	prezime string
	godine  int
	adresa  Adresa
}

func main() {

	osoba1 := Osoba{
		ime:     "Antonija",
		prezime: "Kožul",
		godine:  22,
		adresa: Adresa{
			grad:  "Široki Brijeg",
			ulica: "Knešpolje b.b.",
		},
	}

	fmt.Println()
	osoba1.Ispisi()

}

func (osoba1 Osoba) Ispisi() {
	fmt.Printf("%s %s, %d godine, živi u %s, %s.", osoba1.ime, osoba1.prezime, osoba1.godine, osoba1.adresa.grad, osoba1.adresa.ulica)
}
