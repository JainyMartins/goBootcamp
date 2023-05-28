package exercicio3

import "fmt"

/*Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
Ex: 7 de julho.*/

func Exercicio3() {
	mes := 1

	switch {
	case mes == 1:
		fmt.Println("Janeiro")
	case mes == 2:
		fmt.Println("Fevereiro")
	case mes == 3:
		fmt.Println("Março")
	case mes == 4:
		fmt.Println("Abril")
	case mes == 5:
		fmt.Println("Maio")
	case mes == 6:
		fmt.Println("Junho")
	case mes == 7:
		fmt.Println("Julho")
	case mes == 8:
		fmt.Println("Agosto")
	case mes == 9:
		fmt.Println("Setembro")
	case mes == 10:
		fmt.Println("Outubro")
	case mes == 11:
		fmt.Println("Novembro")
	case mes == 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês desconhecido!")
	}

	//Sim, com switch ou if...else Case porque fica mais clean code.
}
