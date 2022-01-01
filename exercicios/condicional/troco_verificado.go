package main

import "fmt"

/*

Problema "troco_verificado"
	Fazer um programa para calcular o troco no processo de pagamento de um produto de uma mercearia.
	O programa deve ler o preço unitário do produto, a quantidade de unidades compradas deste produto,
	e o valor em dinheiro dado pelo cliente. Seu programa deve mostrar o valor do troco a ser devolvido
	ao cliente. Se o dinheiro dado pelo cliente não for suficiente, mostrar uma mensagem informando o
	valor restante conforme exemplo.

Exemplo 1:
	Preço unitário do produto: 8.00
	Quantidade comprada: 2
	Dinheiro recebido: 20.00
	TROCO = 4.0


*/

func main() {
	var preco, quantidade, dinheiro, troco, total float64
	fmt.Print("Preco unitario do produto : ")
	fmt.Scan(&preco)
	fmt.Print("Quantidade comprada: ")
	fmt.Scan(&quantidade)

	fmt.Print("Dinheiro recebido: ")
	fmt.Scan(&dinheiro)

	total = preco * quantidade

	if total > dinheiro {
		trocoFalta := total - dinheiro
		fmt.Printf("DINHEIRO INSUFICIENTE. FALTAM %.2f REAIS ", trocoFalta)
	} else {
		troco = dinheiro - total
		fmt.Printf("TROCO = %.2f", troco)
	}
}
