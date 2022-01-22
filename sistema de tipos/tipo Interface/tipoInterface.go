package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "opa"
	fmt.Println(coisa2)

	// por ser uma interface vázia, ele se comporto da forma "dinâmica"
	coisa2 = true

	coisa2 = curso{"GoLang: Linguagem do google"}
	fmt.Println(coisa2)
}
