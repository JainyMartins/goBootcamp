package exercicio2

import "fmt"

/* Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
pressão atmosférica de diferentes lugares.
1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?*/

func Exercicio2() {
	var temp int = 26
	var umidade int = 56
	var atm float32 = 1013.25

	fmt.Printf("A temperatura está em %dºC, a umidade em %d%% e a pressão atmosférica em %.2f\n", temp, umidade, atm)
}
