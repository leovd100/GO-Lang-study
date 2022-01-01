package main

import (
	"fmt"
	"math"
)

/*
Problema "baskara"
	Fazer um programa para ler os três coeficientes de uma equação do segundo grau. Usando a fórmula
	de Baskara, calcular e mostrar os valores das raízes x1 e x2 da equação com quatro casas decimais,
	conforme exemplo. Se a equação não possuir raízes reais, mostrar uma mensagem.

Exemplo 1:
	Coeficiente a: 1
	Coeficiente b: 0
	Coeficiente c: -9
	X1 = 3.0000
	X2 = -3.0000
*/
func main() {
	var a, b, c float64

	fmt.Print("Coeficiente a: ")
	fmt.Scanln(&a)
	fmt.Print("Coeficiente b: ")
	fmt.Scanln(&b)
	fmt.Print("Coeficiente c: ")
	fmt.Scanln(&c)

	delta := math.Pow(b, 2) - (4 * (a) * (c))

	if delta < 0 {
		fmt.Print("Esta equacao nao possui raizes reais ")
	} else {
		x1 := ((-b) - math.Sqrt(delta)) / (2 * a)
		x2 := ((-b) + math.Sqrt(delta)) / (2 * a)

		fmt.Printf("x1 = %.4f\n", x1)
		fmt.Printf("X2 = %.4f\n", x2)
	}

}
