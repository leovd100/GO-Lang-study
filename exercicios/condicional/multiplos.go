package main

import "fmt"

/*

Problema "multiplos" (adaptado de URI 1044)
Fazer um programa para ler dois números inteiros, e dizer se um número é múltiplo do outro. Os
números podem ser digitados em qualquer ordem.

Exemplo 1:
Digite dois numeros inteiros:
6
24
Sao multiplos

Exemplo 2:
Digite dois numeros inteiros:
24
6
Sao multiplos

Exemplo 3:
Digite dois numeros inteiros:
13
5
Nao sao multiplos

*/

func main() {
	var v1, v2 int
	fmt.Println("Digite dois numeros inteiros: ")
	fmt.Print("Primeiro: ")
	fmt.Scanln(&v1)
	fmt.Print("Segundo: ")
	fmt.Scanln(&v2)

	if v1 > v2 {
		v1, v2 = v2, v1
	}

	if v1%v2 == v1 && v2%v1 == 0 {
		fmt.Print("Sao Multiplos")
	} else {
		fmt.Print("Nao São Multiplos")
	}

}
