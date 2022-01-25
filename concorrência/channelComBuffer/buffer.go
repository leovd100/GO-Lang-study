package main

import (
	"fmt"
	"time"
)

// O buffer só vai bloquear o envio de mesagem quando ele lotar o buffer
func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou!") // Dependendo da posição em que o Print está, ele pode ser impresso antes do buffer encher
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
