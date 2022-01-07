package main

import "fmt"

/*


Problema "aumento" (adaptado de URI 1048)
Uma empresa vai conceder um aumento percentual de
salário aos seus funcionários dependendo de quanto
cada pessoa ganha, conforme tabela ao lado. Fazer um
programa para ler o salário de uma pessoa, daí mostrar
qual o novo salário desta pessoa depois do aumento,
quanto foi o aumento e qual foi a porcentagem de
aumento


Salário atual  				Aumento
Até R$ 1000.00  			20%
Acima de R$ 1000.00
até R$ 3000.00 				15%

Acima de R$ 3000.00
até R$ 8000.00 				10%

Acima de R$ 8000.00  		5%

Exemplo 1:
Digite o salario da pessoa: 2500.00
Novo salario = R$ 2875.00
Aumento = R$ 375.00
Porcentagem = 15 %

Exemplo 2:
Digite o salario da pessoa: 8000.00
Novo salario = R$ 8800.00
Aumento = R$ 800.00
Porcentagem = 10 %


*/

func main() {
	var salario, novoSalario, aumento float64
	var porcentagem int
	fmt.Print("Digite o salario da pessoa: ")
	fmt.Scanln(&salario)

	switch {

	case salario <= 1000:
		porcentagem, novoSalario, aumento = calculo(20, salario)
	case salario <= 3000:
		porcentagem, novoSalario, aumento = calculo(15, salario)
	case salario <= 8000:
		porcentagem, novoSalario, aumento = calculo(10, salario)
	case salario > 8000:
		porcentagem, novoSalario, aumento = calculo(5, salario)
	default:
		fmt.Printf("Valor inválido")
	}

	fmt.Printf("Novo salario = R$ %.2f\n", novoSalario)
	fmt.Printf("Aumento = R$ %.2f\n", aumento)
	fmt.Printf("Porcentagem = %d Porcento", porcentagem)

}

func calculo(porcentagem int, salario float64) (int, float64, float64) {

	novoSalario := salario + ((float64(porcentagem) / 100) * salario)
	aumento := novoSalario - salario
	return porcentagem, novoSalario, aumento
}
