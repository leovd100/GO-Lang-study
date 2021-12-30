package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Problema "idades"
	Fazer um programa para ler o nome e idade de duas pessoas. Ao final mostrar uma mensagem com os
	nomes e a idade média entre essas pessoas, com uma casa decimal, conforme exemplo.

	Exemplo:
	Dados da primeira pessoa:
	Nome: Maria Silva
	Idade: 19
	Dados da segunda pessoa:
	Nome: Joao Melo
	Idade 20
	A idade média de Maria Silva e Joao Melo é de 19.5 anos

*/

func main() {
	var idade1, idade2 int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Dados da primeira pessoa")
	fmt.Print("Nome: ")
	nome, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("Nome inválido")
	}
	fmt.Print("Idade: ")
	fmt.Scanln(&idade1)

	fmt.Println("Dados da segunda pessoa")
	fmt.Print("Nome: ")

	nome2, err := reader.ReadString('\n')
	fmt.Print("Idade: ")
	fmt.Scanln(&idade2)

	if err != nil {
		fmt.Print("Nome inválido")
	}

	media := float64((idade1 + idade2)) / 2.0

	fmt.Printf("A idade media entre %s e %s e de %.1f", nome, nome2, media)

}
