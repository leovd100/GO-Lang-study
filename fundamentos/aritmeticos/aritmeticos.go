package main

import (
	"fmt"
	"math"
)

func maing() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)

	// bitwise

	fmt.Println("E => ", a&b)
	fmt.Println("OU => ", a|b)
	fmt.Println("Xor => ", a^b)

	c := 3.0
	d := 2.0

	// outras operacoes usando math
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Exponencial => ", math.Pow(c, d))

}
