package main

import "fmt"

func Somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := Somar(5, 8)
	imprimir(resultado)
}
