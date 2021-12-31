package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Problema "pagamento"
	Fazer um programa para ler o nome de um(a) funcionário(a), o valor que ele(a) recebe por hora, e a
	quantidade de horas trabalhadas por ele(a). Ao final, mostrar o valor do pagamento do funcionário com
	uma mensagem explicativa, conforme exemplo.

Exemplo 1:
	Nome: Joao Silva
	Valor por hora: 50.00
	Horas trabalhadas: 60
	O pagamento para Joao Silva deve ser 3000.00

*/

func main() {

	var valorPorHora, total float64
	var trb int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nome: ")

	nome, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("Erro ao digitar o nome")
	}

	fmt.Print("Valor por hora: ")
	fmt.Scanln(&trb)

	fmt.Print("Horas trabalhadas: ")
	fmt.Scanf("%f", &valorPorHora)
	total = valorPorHora * float64(trb)
	fmt.Printf("O pagamento para %s deve ser %.2f", nome, total)

}
