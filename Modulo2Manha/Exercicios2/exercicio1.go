package main

import "fmt"

func main() {
	//Funcao retorne imposto de um salario, 50000 reais - descontado 17%, 150000, 27%
	fmt.Println(calcularImposto(10000))
}

func calcularImposto(salario float64) float64 {
	if salario < 150000 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}
}
