package exercicio4

import "fmt"

/*Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
pediu a ajuda de um desenvolvedor experiente que pode:
● Revisar o código e realizar as devidas correções.*/

func Exercicio4() {
	var sobrenome string = "Silva" //correto
	var idade int = 25             //int não deve estar entre parênteses
	boolean := false               //boolean não deve estar entre parênteses
	var salario float32 = 4585.90  //salario deve ser em float
	var nome string = "Fellipe"    //correto

	fmt.Println(sobrenome, idade, boolean, salario, nome)

}
