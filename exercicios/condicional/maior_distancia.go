package main

import (
	"fmt"
	"math"
)

/*
Problema "dardo"
No arremesso de dardo, o atleta tem três chances para lançar o dardo à maior distância que conseguir.
Você deve criar um programa para, dadas as medidas das três tentativas de lançamento, informar qual
foi a maior.

Exemplo 1:
	Digite as tres distancias:
	83.21
	79.53
	89.15
	MAIOR DISTANCIA = 89.15

Exemplo 2:
	Digite as tres distancias:
	83.21
	87.20
	83.21
	MAIOR DISTANCIA = 87.20


*/

func main() {
	var d1, d2, d3 float64
	fmt.Println("Digite as tres distancias: ")
	fmt.Scanln(&d1)
	fmt.Scanln(&d2)
	fmt.Scanln(&d3)
	resultado := verifica(d1, d2, d3)
	fmt.Println("MAIO DISTANCIA = ", resultado)
}

func verifica(a float64, b float64, c float64) float64 {
	primeiro := math.Max(a, b)
	second := math.Max(primeiro, c)
	return second
}
