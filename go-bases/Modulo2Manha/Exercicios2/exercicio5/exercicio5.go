package exercicio5

import (
	"errors"
	"fmt"
)

/*Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
com o animal especificado e que retorne uma função com uma mensagem (caso não
exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na
quantidade necessária do animal digitado.
Exemplo:*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func caoFunc(qtd int) float64 {
	return 10.0 * float64(qtd)
}

func catFunc(qtd int) float64 {
	return 5.0 * float64(qtd)
}

func hamsterFunc(qtd int) float64 {
	return 0.25 * float64(qtd)
}

func tarantulaFunc(qtd int) float64 {
	return 0.15 * float64(qtd)
}

func verificarAnimal(tipoAnimal string) (func(int) float64, error) {
	switch tipoAnimal {
	case dog:
		return caoFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		return nil, errors.New("O animal digitado não existe!")
	}
}

func verificarErro(err error) {
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
}

func Exercicio5() {
	animalDog, err := verificarAnimal(dog)
	verificarErro(err)

	animalCat, err := verificarAnimal(cat)
	verificarErro(err)

	animalHamster, err := verificarAnimal(hamster)
	verificarErro(err)

	animalTarantula, err := verificarAnimal(tarantula)
	verificarErro(err)

	var amount float64
	amount += animalDog(5)
	amount += animalCat(2)
	amount += animalHamster(4)
	amount += animalTarantula(4)

	fmt.Printf("O total de comida que o abrigo vai precisar será: %.2fkg", amount)
}
