package exercicio1

import (
	"fmt"
	"os"
)

/*1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
tipo “int”.
2. Crie um erro personalizado com uma struct que implemente “Error()” com a
mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
console a mensagem “Necessário pagamento de imposto”.*/
type myError struct{
	msg string
}

func (m *myError) Error() string {
	return fmt.Sprintf("%v", m.msg)
}

func myErrorTest (salario int) (string, error){
	if salario < 15000 {
		return "", &myError{
			msg: "erro: o salário digitado não está dentro do valor mínimo",
		}
	}
	return "Necessário pagamento de imposto", nil
}

func Exercicio1(){
	var salario = 15000
	retorno, erro := myErrorTest(salario)
	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}
	fmt.Println(retorno)
}

