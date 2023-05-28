package exercicio3

import "fmt"

/*Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
de horas trabalhadas por mês e na categoria do funcionário.
Se a categoria for C, seu salário é de R$1.000 por hora
Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
mensais
Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais*/

func Exercicio3() {
	//C - 1000/h
	//B - 1500/h + 20% se mais de 160h mensais
	//A - 3000/h + 50% se mais de 160h mensais
	fmt.Printf("%.2f\n", calculaSalario("C", 162))
	fmt.Printf("%.2f\n", calculaSalario("B", 176))
	fmt.Printf("%.2f\n", calculaSalario("A", 172))
}

func calculaSalario(categoria string, horas int) float64 {
	salario := 0.0
	if categoria == "C" {
		salario = 1000.0
		return salario
	} else if categoria == "B" {
		salario = 1500.0
		if horas > 160 {
			salario *= 1.2
			return salario
		}
		return salario
	} else if categoria == "A" {
		salario = 3000.0
		if horas > 160 {
			salario *= 1.5
		}
		return salario
	} else {
		return 0
	}
}
