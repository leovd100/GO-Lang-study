package main

import "fmt"

/*
Problema "menor_de_tres"
	Fazer um programa para ler três números inteiros. Em seguida, mostrar qual o menor dentre os três
	números lidos. Em caso de empate, mostrar apenas uma vez.

Exemplo 1:
	Primeiro valor: 7
	Segundo valor: 3
	Terceiro valor: 8
	MENOR = 3
*/
func main() {
	var a, b, c int

	fmt.Print("Primeiro valor: ")
	fmt.Scan(&a)
	fmt.Print("Segundo valor: ")
	fmt.Scan(&b)
	fmt.Print("Terceiro valor: ")
	fmt.Scan(&c)

	if a < b && a < c {
		fmt.Print("MENOR = ", a)
	} else if b < c {
		fmt.Print("MENOR = ", b)
	} else {
		fmt.Print("MENOR = ", c)
	}

}
