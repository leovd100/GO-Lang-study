package main

import "fmt"

/*
Problema "duracao"
Fazer um programa para ler uma duração de tempo em segundos, daí imprimir na tela esta duração no
formato horas:minutos:segundos.

Exemplo 1:
	Digite a duracao em segundos: 300
	0:5:0

Exemplo 2:
	Digite a duracao em segundos: 12506
	3:28:26

Exemplo 3:
	Digite a duracao em segundos: 140811
	39:6:51
*/
func main() {

	var tempo, hora, minuto, segundo int
	fmt.Print("Digite a duracao em segundos: ")
	fmt.Scanln(&tempo)

	hora = (tempo / 60) / 60
	minuto = (tempo / 60) % 60
	segundo = tempo % 60

	fmt.Printf("%d:%d:%d", hora, minuto, segundo)

}
