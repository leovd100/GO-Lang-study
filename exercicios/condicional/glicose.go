package main

import "fmt"

/*
Problema "glicose"
Fazer um programa para ler a quantidade de glicose
no sangue de uma pessoa e depois mostrar na tela a
classificação desta glicose de acordo com a tabela de
referência ao lado.

  Classificação  Glicose
  	Normal  Até 100 mg/dl

	Elevado
   	Maior que 100 até   140 mg/dl

	Diabetes
	Maior de 140 mg/dl

Exemplo 1:
Digite a medida da glicose: 90.0
Classificacao: normal

Exemplo 2:
Digite a medida da glicose: 140.0
Classificacao: elevado

Exemplo 3:
Digite a medida da glicose: 143.2
Classificacao: diabete

*/

func main() {
	var glic float64
	fmt.Print("Digite a medida da glicose: ")
	fmt.Scanln(&glic)

	if glic < 100 {
		fmt.Print("Normal")
	} else if glic >= 100 && glic <= 140 {
		fmt.Print("Elevado")
	} else {
		fmt.Print("Diabetes")
	}
}
