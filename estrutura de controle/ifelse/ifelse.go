package main

import "fmt"

func imprimeirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota: ", nota)
	} else {
		fmt.Println("Reprovado com nota: ", nota)
	}
}

func main() {
	imprimeirResultado(8)
	imprimeirResultado(5.1)
}
