package exercicio2

import "fmt"

/*Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.*/

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
			fmt.Println("Parabéns, você poderá pegar um empréstimo, mas com juros!")
		}
	default:
		fmt.Println("Infelizmente você não consegue pegar um empréstimo!")
	}
}
