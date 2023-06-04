package exercicio1

import (
	"fmt"
	"strings"
)

/*A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
das letras separadamente para soletrá-la. Para isso terão que:
1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
letras que ela contém.
2. Em seguida, imprimi cada uma das letras.*/

func Exercicio1() {
	palavra := "Jainy"

	fmt.Println(len(palavra))

	for _, slice := range strings.Split(palavra, "") {
		fmt.Printf("%s ", slice)
	}
}
