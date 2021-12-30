package main

import "fmt"

func main() {
	var x int64 = 1
	y := 2

	//apenas posfix

	x++ //x +=1 ou x + 1
	fmt.Println(x)

	y--
	fmt.Println(y)

	// Não permitido
	// fmt.Println(x == y--)
}
