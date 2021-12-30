package main

import (
	"fmt"
)

/*
Problema "terreno"
Fazer um programa para ler as medidas da largura e comprimento de um terreno retangular com uma
casa decimal, bem como o valor do metro quadrado do terreno com duas casas decimais. Em seguida,
o programa deve mostrar o valor da área do terreno, bem como o valor do preço do terreno, ambos com
duas casas decimais, conforme exemplo.

Exemplo 1:
Digite a largura do terreno: 10.0
Digite o comprimento do terreno: 30.0
Digite o valor do metro quadrado: 200.00
Area do terreno = 300.00
Preco do terreno = 60000.00

*/

func main() {
	var largura, comprimento, valor float64

	fmt.Print("Digite a largura do terreno: ")
	fmt.Scanln(&largura)

	fmt.Print("Digite o comprimento do terreno: ")
	fmt.Scanln(&comprimento)

	fmt.Print("Digite o valor do metro quadrado: ")
	fmt.Scanln(&valor)

	area := comprimento * largura
	total := area * valor

	fmt.Printf("Area do terreno = %.2f\n", area)
	fmt.Printf("Preço do terreno = %.2f", total)
}
