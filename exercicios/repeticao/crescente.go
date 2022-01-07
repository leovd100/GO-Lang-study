package main

import "fmt"

func main() {
	var x, y int
	x, y = menu()
	for x != y {
		x, y = menu()
		verifica(x, y)
	}

}

func menu() (int, int) {
	var x, y int
	fmt.Println("Digite outros dois numeros:")
	fmt.Print("Primeiro valor: ")
	fmt.Scanln(&x)
	fmt.Print("Segundo valor: ")
	fmt.Scanln(&y)
	return x, y
}

func verifica(x, y int) {
	if x > y && x != y {
		fmt.Println("CRESCENTE")
	} else if x != y {
		fmt.Println("DECRESCENTE")
	}
}
