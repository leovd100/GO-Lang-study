package main

import "fmt"

/*

Problema "soma"
	Fazer um programa para ler dois valores inteiros X e Y, e depois mostrar na tela o valor da soma destes
	números.

Exemplo 1:
	Digite o valor de X: 8
	Digite o valor de Y: 10
	SOMA = 18

Exemplo 2:
	Digite o valor de X: 12
	Digite o valor de Y: 31
	SOMA = 43

*/

func main() {
	var x, y int

	fmt.Print("Digite o valor de X: ")
	fmt.Scanln(&x)
	fmt.Print("Digite o valor de Y: ")
	fmt.Scanln(&y)

	resultado := soma(x, y)

	fmt.Printf("SOMA = %d", resultado)
}

func soma(x, y int) int {
	return x + y
}
