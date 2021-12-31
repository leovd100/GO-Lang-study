package main

import "fmt"

/*
Problema "notas"
	Fazer um programa para ler as duas notas que um aluno obteve no primeiro e segundo semestres de
	uma disciplina anual. Em seguida, mostrar a nota final que o aluno obteve (com uma casa decimal) no
	ano juntamente com um texto explicativo. Caso a nota final do aluno seja inferior a 60.00, mostrar a
	mensagem "REPROVADO", conforme exemplos.

Exemplo 1:
	Digite a primeira nota: 45.5
	Digite a segunda nota: 31.3
	NOTA FINAL = 76.8

Exemplo 2:
	Digite a primeira nota: 34.0
	Digite a segunda nota: 23.5
	NOTA FINAL = 57.5
	REPROVADO
*/

func main() {

	var nota1, nota2, notaFinal float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scanln(&nota1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scanln(&nota2)

	notaFinal = nota1 + nota2
	fmt.Println("NOTA FINAL = ", notaFinal)

	if notaFinal < 60.00 {
		fmt.Print("REPROVADO")
	}

}
