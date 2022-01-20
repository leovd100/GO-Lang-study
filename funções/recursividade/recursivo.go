package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatoralAnterior, _ := fatorial(n - 1)
		return n * fatoralAnterior, nil
	}

}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)
	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
