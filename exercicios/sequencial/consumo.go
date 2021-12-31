package main

import "fmt"

/*

Problema "consumo"
	Fazer um programa para ler a distância total (em km) percorrida por um carro, bem como o total de
	combustível gasto por este carro ao percorrer tal distância. Seu programa deve mostrar o consumo
	médio do carro, com três casas decimais.

Exemplo 1:
	Distancia percorrida: 500
	Combustível gasto: 38.5
	Consumo medio = 12.987

*/

func main() {

	var distancia int
	var combustivel, consumo float64

	fmt.Print("Distancia percorrida: ")
	fmt.Scanln(&distancia)

	fmt.Print("Combustivel gasto: ")
	fmt.Scanln(&combustivel)

	consumo = float64(distancia) / combustivel
	fmt.Printf("Consumo medio = %.3f", consumo)
}
