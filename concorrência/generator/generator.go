package main

import (
	"fmt"
	"leovd100/html"
)

// <- chan - canal somente leitura

func main() {
	t1 := html.Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := html.Titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)

}
