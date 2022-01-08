package main

import "fmt"

/*
Problema "soma_vetor"
Faça um programa que leia N números reais e armazene-os em um vetor. Em seguida:
- Imprimir todos os elementos do vetor
- Mostrar na tela a soma e a média dos elementos do vetor

Exemplo:
Quantos numeros voce vai digitar? 4
Digite um numero: 8.0
Digite um numero: 4.0
Digite um numero: 10.0
Digite um numero: 14.0

VALORES = 8.0  4.0  10.0  14.0
SOMA = 36.00
MEDIA = 9.00
*/
func main() {
	var x int
	var soma, media, j float64

	fmt.Print("Quantos numeros voce vai digitar ? ")
	fmt.Scanln(&x)
	slice := make([]float64, 0, x)

	for i := 0; i < x; i++ {
		fmt.Print("Digite um numero: ")
		fmt.Scanln(&j)
		slice = append(slice, j)
		soma += j
	}

	fmt.Println("VALORES = ", slice)
	media = soma / float64(len(slice))
	fmt.Printf("SOMA = %.2f\n", soma)
	fmt.Printf("MEDIA = %.2f\n", media)

}
