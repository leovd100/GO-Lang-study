package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// Mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[97878748484] = "Pedro"
	aprovados[15486353954] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[97878748484]) // ler map
	delete(aprovados, 97878748484)      // deleta a chave
	fmt.Println(aprovados[97878748484])
}
