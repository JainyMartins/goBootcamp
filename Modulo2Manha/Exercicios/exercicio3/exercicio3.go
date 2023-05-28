package exercicio3

import "fmt"

/*Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
2. Corrigir as incorrectas.*/

func Exercicio3() {
	var nome string               //não poderia começar com número (1nome)
	var sobrenome string          //correto
	var idade int                 //o nome da variável vem antes do seu tipo
	var sobrenome2 string         //não devemos usar número no início na declaração, string não é inteiro
	var licencaParaDirigir = true //não devemos usar underlines dessa forma
	var estaturaDaPessoa float32  //não devemos ter espaços entre as palavras da variável
	quantidadeDeFilhos := 2       //correto

	fmt.Println(nome, sobrenome, idade, sobrenome2, licencaParaDirigir, estaturaDaPessoa, quantidadeDeFilhos)
}
