package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método com receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Camisa",
		preco:    29.90,
		desconto: 0.2,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.8}
	fmt.Println(produto2)
}
