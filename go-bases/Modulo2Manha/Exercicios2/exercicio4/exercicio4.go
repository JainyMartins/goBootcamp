package exercicio4

import (
	"errors"
	"fmt"
	"sort"
)

/*Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior
Exemplo:*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
	erro    = "Erro: "
)

func minimumFunc(valores ...float64) float64 {
	sort.Float64s(valores)
	return valores[0]
}

func maximumFunc(valores ...float64) float64 {
	sort.Float64s(valores)
	return valores[len(valores)-1]
}

func averageFunc(valores ...float64) float64 {
	soma := 0.0
	for _, valor := range valores {
		soma += valor
	}
	media := soma / float64(len(valores))
	return media
}

func operation(tipoCalc string) (func(...float64) float64, error) {
	switch tipoCalc {
	case minimum:
		return minimumFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maximumFunc, nil
	default:
		return nil, errors.New("Tipo de Cálculo Inválido!")
	}
}

func verificarErro(err error) {
	if err != nil {
		fmt.Println(erro, err)
		return
	}
}

func Exercicio4() {
	//valor min, max e medio, funcao que indique qual calculo sera realizado - min, max ou media e
	//que retorna outra funcao - e uma mensagem caso o calculo n esteja definido que pode ser
	//passado quantidade N de interios e retorne o calculo que foi indicado na funcao anterior

	//Calculando o valor mínimo
	minFun, err := operation(minimum)
	verificarErro(err)
	min := minFun(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Valor mínimo: %.2f\n", min)

	//Calculando o valor médio
	avgFun, err := operation(average)
	verificarErro(err)
	avg := avgFun(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Valor médio: %.2f\n", avg)

	//Calculando o valor máximo
	maxFun, err := operation(maximum)
	verificarErro(err)
	max := maxFun(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Valor máximo: %.2f\n", max)
}
