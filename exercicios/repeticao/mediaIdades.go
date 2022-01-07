package main

import "fmt"

/*
Problema "media_idades"
Faça um programa para ler um número indeterminado de dados, contendo cada um, a idade de um
indivíduo. O último dado, que não entrará nos cálculos, contém um valor de idade negativa. Calcular
e imprimir a idade média deste grupo de indivíduos. Se for entrado um valor negativo na primeira vez,
mostrar a mensagem "IMPOSSIVEL CALCULAR".

Exemplo 1:
Digite as idades:
31
27
46
-5
MEDIA = 34.67

Exemplo 2:
Digite as idades:
-10
IMPOSSIVEL CALCULAR
*/
func main() {
	var somaIdade, i, idade int
	fmt.Println("Digite as idades: ")
	fmt.Scanln(&idade)
	i, somaIdade = verifica(idade, i, somaIdade)
	for idade > 0 {
		fmt.Scanln(&idade)
		i, somaIdade = verifica(idade, i, somaIdade)
	}

	if somaIdade <= 0 {
		fmt.Print("IMPOSSIVEL CALCULAR")
	} else {
		media := float64(somaIdade) / float64(i)
		fmt.Printf("MEDIA = %.2f\n", media)
	}
}

func verifica(x, y, z int) (int, int) {
	if x > 0 {
		y++
		z += x
	}
	return y, z
}
