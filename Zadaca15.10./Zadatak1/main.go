// 	Napiši strukturu koja predstavlja pravokutnik i sadrži polja za širinu i visinu.
// 	Kreiraj metode za izračunavanje površine i opsega pravokutnika,
// 	metode moraju biti vezane za novo kreiranu strukturu Pravokutnika.
// 	U main funkciji inicijalizirajte jedan pravokutnik, te pozovite iznad kreirane metode za računanje
// 	površine i opsega.

package main

import "fmt"

type Pravokutnik struct {
	sirina float32
	visina float32
}

func main() {

	pravokutnik1 := Pravokutnik{
		sirina: 2.5,
		visina: 3.5,
	}

	pravokutnik1.izracunajPovrsinu()
	fmt.Println()
	pravokutnik1.izracunajOpseg()
}

func (pravokutnik1 Pravokutnik) izracunajPovrsinu() {
	var Povrsina float32
	Povrsina = pravokutnik1.sirina * pravokutnik1.visina

	fmt.Printf("Povrsina pravokutnika je %f cm.", Povrsina)
}

func (pravokutnik1 Pravokutnik) izracunajOpseg() {
	var Opseg float32
	Opseg = 2*pravokutnik1.sirina + 2*pravokutnik1.visina

	fmt.Printf("Opseg pravokutnika je %f cm.", Opseg)
}
