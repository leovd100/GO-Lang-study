package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só depois que for lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c)

	fmt.Println("Foi lido")
	fmt.Print(<-c) //deadlock
	fmt.Println("Fim")
}
