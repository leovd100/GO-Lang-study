package main

import "fmt"

func main() {
	fmt.Print("Mesa")
	fmt.Print(" linha.")

	fmt.Println(" Nova linha")
	fmt.Println("Linha.")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é  = " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %t %s", a, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
