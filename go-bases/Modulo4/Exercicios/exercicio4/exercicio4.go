package exercicio4

import (
	"errors"
	"fmt"
)

/*Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
1. Desenvolva as funções necessárias para permitir que a empresa calcule:
a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
- A função receberá as horas trabalhadas no mês e o valor da hora como
parâmetro.
- Esta função deve retornar mais de um valor (salário calculado e erro).
- No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
descontados como imposto.

2

- Se o número de horas mensais inseridas for menor que 80 ou um número
negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
não pode ter trabalhado menos de 80 horas por mês".

2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando
“errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos
retornos de erro em sua função “main()”.*/

func calcularSalario(horasMensais, valorHora float64) (float64, error) {
	salario := horasMensais * valorHora
	if salario >= 20000.00 {
		salario *= 0.9
	}
	if horasMensais < 80 || horasMensais < 0 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}
	return salario, nil
}

func Exercicio4() {
	salario, erro := calcularSalario(80, 300)
	if erro != nil {
		fmt.Println(erro)
		return
	}
	fmt.Printf("O salário mensal do trabalhador é: R$ %.2f", salario)
}
