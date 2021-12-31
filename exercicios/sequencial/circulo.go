package main

import (
	"fmt"
	"math"
)

/*
Problema "circulo"
	Fazer um programa para ler o valor "r" do raio de um círculo, e depois mostrar o valor da área do
	círculo com três casas decimais. A fórmula da área do círculo é a seguinte: 𝑎𝑟𝑒𝑎=𝜋.𝑟ଶ. Você pode
	usar o valor de 𝜋 fornecido pela biblioteca da sua linguagem de programação, ou então, se preferir, use
	diretamente o valor 3.14159.

Exemplo 1:
	Digite o valor do raio do circulo: 2.0
	AREA = 12.566

*/
func main() {
	pi := 3.14159

	var area, raio float64

	fmt.Print("Digite o valor do raio do circulo: ")
	fmt.Scanln(&raio)
	area = pi * math.Pow(raio, 2)
	fmt.Println("Area = ", area)
}
