package exercicio2

import "fmt"

func Exercicio2() {
	idade := 23
	situacaoTrabalho := "Empregado"
	anosTrabalho := 2
	salario := 150001

	switch {
	case idade > 22 && situacaoTrabalho == "Empregado" && anosTrabalho > 1:
		switch {
		case salario > 150000:
			fmt.Println("Parabéns, você poderá pegar um empréstimo sem juros!")
		default:
			fmt.Println("Parabés, você poderá pegar um empréstimo, mas com juros!")
		}
	default:
		fmt.Println("Infelizmente você não consegue pegar um empréstimo!")
	}
}
