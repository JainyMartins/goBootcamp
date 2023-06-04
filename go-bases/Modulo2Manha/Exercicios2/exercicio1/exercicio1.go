package exercicio1

import "fmt"

/*Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.*/

func Exercicio1() {
	//Funcao retorne imposto de um salario, 50000 reais - descontado 17%, 150000, 27%
	fmt.Printf("O imposto será no valor de: %.2f\n", calcularImposto(10000))
}

func calcularImposto(salario float64) float64 {
	if salario < 150000 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}
}
