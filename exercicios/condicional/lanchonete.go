package main

import "fmt"

/*

Problema "lanchonete" (adaptado de URI 1038)
Uma lanchonete possui vários produtos. Cada produto possui um código
e  um  preço.  Você  deve  fazer  um  programa  para  ler  o  código  e  a
quantidade comprada de um produto (suponha um código válido), e daí
informar qual o valor a ser pago, com duas casas decimais, conforme
tabela de produtos ao lado.


Exemplo 1:
Codigo do produto comprado: 1
Quantidade comprada: 3
Valor a pagar: R$ 15.00

Exemplo 2:
Codigo do produto comprado: 4
Quantidade comprada: 2
Valor a pagar: R$ 17.80

Código do
produto
			Preço do
			produto
1  			R$ 5.00
2  			R$ 3.50
3  			R$ 4.80
4  			R$ 8.90
5  			R$ 7.32

*/

func main() {
	var cd, qnt int
	var preco float64
	fmt.Print("Codigo do produto comprado: ")
	fmt.Scanln(&cd)
	fmt.Print("Quantidade comprada: ")
	fmt.Scanln(&qnt)
	if cd == 1 {
		preco = 5.00 * float64(qnt)
	} else if cd == 2 {
		preco = 3.50 * float64(qnt)
	} else if cd == 3 {
		preco = 4.80 * float64(qnt)
	} else if cd == 4 {
		preco = 8.90 * float64(qnt)
	} else if cd == 5 {
		preco = 7.32 * float64(qnt)
	}
	fmt.Printf("Valor a pagar: %.2f\n", preco)
}
