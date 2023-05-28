package main

import "fmt"

func main() {
	//C - 1000/h
	//B - 1500/h + 20% se mais de 160h mensais
	//A - 3000/h + 50% se mais de 160h mensais
	fmt.Printf("%.2f\n", calculaSalario("C", 162))
	fmt.Printf("%.2f\n", calculaSalario("B", 176))
	fmt.Printf("%.2f\n", calculaSalario("A", 172))
}

func calculaSalario(categoria string, horas int) float64 {
	salario := 0.0
	if categoria == "C" {
		salario = 1000.0
		return salario
	} else if categoria == "B" {
		salario = 1500.0
		if horas > 160 {
			salario *= 1.2
			return salario
		}
		return salario
	} else if categoria == "A" {
		salario = 3000.0
		if horas > 160 {
			salario *= 1.5
		}
		return salario
	} else {
		return 0
	}
}
