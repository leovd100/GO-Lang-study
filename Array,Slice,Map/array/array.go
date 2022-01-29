package main

import "fmt"

func main() {
	//homogênea (mesmo tipo de dado) e estático (fixo)

	var nota [3]float64
	fmt.Println(nota)

	nota[0], nota[1], nota[2] = 7.8, 4.3, 9.1
	fmt.Println(nota)

	total := 0.0
	for i := 0; i < len(nota); i++ {
		total += nota[i]
	}

	media := total / float64(len(nota))
	fmt.Printf("Media %.2f\n", media)
}
