package main

import "fmt"

func main() {
	ch := make(chan int, 1) // Tipo de dado a ser trafegado no channel e o buffered
	ch <- 1                 // enviando dados para o canal
	<-ch                    //Rece dados do canal
	ch <- 2

	fmt.Println(<-ch)
}
