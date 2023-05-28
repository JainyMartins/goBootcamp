package exercicio2

import (
	"errors"
	"fmt"
)

/*Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro*/

func Exercicio2() {
	//Calcular media das notas por aluno, possamos passar n quantidade de inteiros e devolver media, caso seja negativo, retornar erro
	resultado, erro := calculaMedia(2, 2, 4, 5, 10)
	if erro != nil {
		fmt.Println("Houve um erro, resultado: ", resultado, erro.Error())
	} else {
		fmt.Println("O resultado é:", resultado)
	}
}

func calculaMedia(notas ...int) (float64, error) {
	soma := 0
	for _, valor := range notas {
		if valor < 0 {
			return 0, errors.New("Existe uma nota negativa!")
		}
		soma += valor
	}
	return float64(soma / len(notas)), nil
}
