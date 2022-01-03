package main

import "fmt"

/*

Problema "tempo_de_jogo" (adaptado de URI 1046)
Leia a hora inicial e a hora final de um jogo. A seguir calcule a duração do jogo, sabendo que o mesmo
pode começar em um dia e terminar em outro, tendo uma duração mínima de 1 hora e máxima de 24
horas.

Exemplo 1:
Hora inicial: 16
Hora final: 2
O JOGO DUROU 10 HORA(S)

Exemplo 2:
Hora inicial: 0
Hora final: 0
O JOGO DUROU 24 HORA(S)

Exemplo 3:
Hora inicial: 2
Hora final: 16
O JOGO DUROU 14 HORA(S)


*/

func main() {
	var inicio, final, duracao int
	fmt.Print("Hora inicial: ")
	fmt.Scanln(&inicio)
	fmt.Print("Hora final: ")
	fmt.Scanln(&final)

	if inicio > final {
		duracao = 24 - (inicio - final)
	} else if inicio == final {
		duracao = 24
	} else {
		duracao = final - inicio
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)", duracao)
}
