package main

import "fmt"

/*
Problema "medidas"
	Fazer um programa para ler três medidas A, B e C. Em seguida, calcular e mostrar (imprimir os dados
	com quatro casas decimais):
	a) a área do quadrado que tem lado A
	b) a área do triângulo retângulo que base A e altura B
	c) a área do trapézio que tem bases A e B, e altura C

Exemplo 1:
	Digite a medida A: 4.0
	Digite a medida B: 3.5
	Digite a medida C: 5.2
	AREA DO QUADRADO = 16.0000
	AREA DO TRIANGULO = 7.0000
	AREA DO TRAPEZIO = 19.5000

*/

func main() {
	var a, b, c, quadrado, triangulo, trapezio float64

	fmt.Print("Digite a medida A: ")
	fmt.Scanln(&a)
	fmt.Print("Digite a medida B: ")
	fmt.Scanln(&b)
	fmt.Print("Digite a medida C: ")
	fmt.Scanln(&c)

	quadrado = a * a
	triangulo = a * b
	trapezio = ((a + b) * c) / 2

	fmt.Printf("AREA DO QUADRADO = %.4f\n", quadrado)
	fmt.Printf("AREA DO TRIANGULO = %.4f\n", triangulo)
	fmt.Printf("AREA DO TRAPEZIO = %.4f", trapezio)

}
