package main

import (
	"fmt"
	"strings"
)

/*
Problema "temperatura"
Deseja-se converter uma medida de temperatura da escala Celsius para Fahrenheit ou vice-versa. Para
isso, você deve construir um programa que leia a letra "C" ou "F" indicando em qual escala vai ser
informada uma temperatura. Em seguida o programa deve mostrar a temperatura na outra escala com
duas casas decimais. A seguir é dada a fórmula para converter de Fahrenheit para Celsius (você deve
deduzir a fórmula de Celsius para Fahrenheit):  )32(9
5  FC

Exemplo 1:
Voce vai digitar a temperatura em qual escala (C/F)? F
Digite a temperatura em Fahrenheit: 75.00
Temperatura equivalente em Celsius: 23.89

Exemplo 2:
Voce vai digitar a temperatura em qual escala (C/F)? C
Digite a temperatura em Celsius: 28.15
Temperatura equivalente em Fahrenheit: 82.67

*/

func main() {
	var celsius, fahrenheit float64
	var char string
	fmt.Print("Voce vai digitar a temperatura em qual escala (C/F)?")
	fmt.Scanln(&char)
	low := strings.ToLower(char)
	if low == "f" {

		fmt.Print("Digite a temperatura em Fahrenheit: ")
		fmt.Scanln(&fahrenheit)
		resultado := (float64(5) / float64(9)) * (fahrenheit - 32)
		fmt.Printf("Temperatura equivalente em Celsius: %.2f", resultado)
	} else {
		fmt.Print("Digite a temperatura em Celsius: ")
		fmt.Scanln(&celsius)
		resultado := (9/5)*celsius + 32
		fmt.Printf("Temperatura equivalente em Fahrenheit: %.2f", resultado)
	}

}
