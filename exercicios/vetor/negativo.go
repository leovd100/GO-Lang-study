package main

import "fmt"

/*
Problema "negativos"
Faça um programa que leia um número inteiro positivo N (máximo = 10) e depois N números inteiros
e armazene-os em um vetor. Em seguida, mostrar na tela todos os números negativos lidos.

Exemplo:
Quantos numeros voce vai digitar? 6
Digite um numero: 8
Digite um numero: -2
Digite um numero: 9
Digite um numero: 10
Digite um numero: -3
Digite um numero: -7
NUMEROS NEGATIVOS:
-2
-3
-7
*/

func main() {
	var j, x int
	fmt.Print("Quanto numeros você vai digitar ? ")
	fmt.Scanln(&j)

	s := make([]int, 1, j)

	for i := 0; i < j; i++ {
		fmt.Print("Digite um numero : ")
		fmt.Scanln(&x)
		s = append(s, x)
	}

	for _, valor := range s {

		if valor < 0 {
			fmt.Println(valor)
		}
	}

}
