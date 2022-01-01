package main

import "fmt"

/*
Problema "operadora"
	Uma operadora de telefonia cobra R$ 50.00 por um plano básico que dá direito a 100 minutos de
	telefone. Cada minuto que exceder a franquia de 100 minutos custa R$ 2.00. Fazer um programa para
	ler a quantidade de minutos que uma pessoa consumiu, daí mostrar o valor a ser pago.

Exemplo 1:
	Digite a quantidade de minutos: 22
	Valor a pagar: R$ 50.00

*/

func main() {
	valorInicial := 50.00
	var excedente float64
	var minutos int

	fmt.Print("Digite a quantidade de minutos : ")
	fmt.Scan(&minutos)

	if minutos > 100 {
		excedente = (float64(minutos) - 100) * 2.0
	}
	valorInicial += excedente
	fmt.Printf("Valor a pagar: %.2f\n", valorInicial)
}
