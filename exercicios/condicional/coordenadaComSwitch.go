package main

import "fmt"

/*
Exemplo 1:
Valor de X: 4.5
Valor de Y: -2.2
Q4

Exemplo 2:
Valor de X: 3.1
Valor de Y: 2.0
Q1

Exemplo 3:
Valor de X: 0
Valor de Y: 0
Origem

Exemplo 4:
Valor de X: 3.8
Valor de Y: 0
Eixo X

*/
func main() {
	var x, y float64

	fmt.Print("Valor de X: ")
	fmt.Scanln(&x)
	fmt.Print("Valor de y: ")
	fmt.Scanln(&y)

	resultado := Calculo(x, y)
	fmt.Print(resultado)
}

func Calculo(x float64, y float64) string {
	switch {
	case x == 0 && y == 0:
		return "ORIGEM"
	case x == 0 && y > 0:
		return "EIXO Y"
	case x > 0 && y == 0:
		return "EIXO X"
	case x > 0 && y > 0:
		return "Q1"
	case x < 0 && y > 0:
		return "Q2"
	case x < 0 && y < 0:
		return "Q3"
	default:
		return "Q4"

	}
}
