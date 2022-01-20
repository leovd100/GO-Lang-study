package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *

func inc2(n *int) {
	// * é usado para desrefenreciar, ou seja, ter acesso ao valor no qual o ponteiro aponta

	*n++
}

func main() {
	n := 1

	inc1(n) // Aqui passa uma cópia da referencia não modificando o original
	// & utilizado para obter um endereço de uma variável
	fmt.Println(n)
	inc2(&n) // Passando por referência
	fmt.Println(n)
}
