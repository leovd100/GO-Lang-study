package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Saudável: %t", tv50, tv32, !sorvete)
}
