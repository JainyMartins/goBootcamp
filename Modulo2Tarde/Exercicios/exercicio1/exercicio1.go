package exercicio1

import (
	"fmt"
	"strings"
)

func Exercicio1() {
	palavra := "Jainy"

	fmt.Println(len(palavra))

	for _, slice := range strings.Split(palavra, "") {
		fmt.Printf("%s ", slice)
	}
}
