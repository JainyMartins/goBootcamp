package exercicio2

import (
	"errors"
	"fmt"
)

/*Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
“Error()”, seja implementado “errors.New()”.*/

func Exercicio2() {
	salario := 10000
	if salario < 15000 {
		fmt.Println(errors.New("erro: o salário digitado não está dentro do valor mínimo"))
		return
	}
	fmt.Println("Necessário pagamento de imposto")
}