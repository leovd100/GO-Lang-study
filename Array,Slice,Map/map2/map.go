package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, // Obrigatório o último elemento ter vírgula
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") // não gera mensagem de erro

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
