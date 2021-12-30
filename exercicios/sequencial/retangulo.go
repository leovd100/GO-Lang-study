package main

import (
	"fmt"
	"math"
)

/*
Problema "retangulo"
	Fazer um programa para ler as medidas da base e altura de um retângulo. Em seguida, mostrar o valor
	da área, perímetro e diagonal deste retângulo, com quatro casas decimais, conforme exemplos.

	Exemplo 1:
	Base do retangulo: 4.0
	Altura do retangulo: 5.0
	AREA = 20.0000
	PERIMETRO = 18.0000
	DIAGONAL = 6.4031

*/

func main() {
	var base, altura, area, perimetro, diagonal float64

	fmt.Print("Base do retangulo: ")
	fmt.Scanln(&base)

	fmt.Print("Altura do retangulo: ")
	fmt.Scanln(&altura)

	/* area = base * altura
	perimetro = (base * 2) + (altura * 2)
	exp := (math.Pow(base, 2) + math.Pow(altura, 2))
	diagonal = math.Sqrt(exp)
	*/

	area, perimetro, diagonal = calc(base, altura)

	fmt.Printf("AREA = %.4f\n", area)
	fmt.Printf("PERIMETRO = %.4f\n", perimetro)
	fmt.Printf("DIAGONAL = %.4f\n", diagonal)

}

func calc(x, y float64) (float64, float64, float64) {

	area := x * y
	perimetro := (x * 2) + (y * 2)
	exp := (math.Pow(x, 2) + math.Pow(y, 2))
	diagonal := math.Sqrt(exp)

	return area, perimetro, diagonal
}
