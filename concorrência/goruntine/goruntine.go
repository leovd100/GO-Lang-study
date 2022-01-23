package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second) // One second

		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	/* 	fale("Leo", "Pq você não fala comigo", 3)
	   	fale("João", "Só posso favar depois de você", 1) */

	// O comando go, cria uma goRuntine executando
	// O código de forma independênte
	// Uma thread pode executar várias goRuntine
	// Uma goRutine só vai continuar executando, se a thread principal do seu programa estiver funcionando
	/* 	go fale("Maria", "Ei...", 500)
	   	go fale("João", "Opa...", 500) */

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns", 5)
}
