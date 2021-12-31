package main

import "fmt"

/*

Problema "troco"

	Fazer um programa para calcular o troco no processo de pagamento de um produto de uma mercearia.
	O programa deve ler o preço unitário do produto, a quantidade de unidades compradas deste produto,
	e o valor em dinheiro dado pelo cliente (suponha que haja dinheiro suficiente). Seu programa deve
	mostrar o valor do troco a ser devolvido ao cliente.

Exemplo 1:
	Preço unitário do produto: 8.00
	Quantidade comprada: 2
	Dinheiro recebido: 20.00
	TROCO = 4.00

Exemplo 2:
	Preço unitário do produto: 30.00
	Quantidade comprada: 3
	Dinheiro recebido: 100.00
	TROCO = 10.00

*/

func main() {
	var quantidade int
	var precoUni, dinheiro, troco float64

	fmt.Print("Preco unitario do produto: ")
	fmt.Scanln(&precoUni)
	fmt.Print("Quantidade comprada: ")
	fmt.Scanln(&quantidade)
	fmt.Print("Dinheiro recebido: ")
	fmt.Scanln(&dinheiro)

	troco = dinheiro - (precoUni * float64(quantidade))
	fmt.Printf("TROCO = %.2f", troco)
}
