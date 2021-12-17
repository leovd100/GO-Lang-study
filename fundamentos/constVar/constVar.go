package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // Tipo [float64] inferido pelo compilador

	// forma reduzida

	area := PI * m.Pow(raio, 2)
	fmt.Print(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Epa!"

	fmt.Println(g, h, i)
}
