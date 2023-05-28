package exercicio1

import (
	"fmt"
	"time"
)

/*Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
detalhes dos dados de cada um deles, conforme o exemplo abaixo:
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
alunos.
Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
Data e que tenha um método de detalhamento*/

type Alunos struct {
	Nome         string
	Sobrenome    string
	RG           string
	DataAdmissao time.Time
}

func (a Alunos) detalhes() {
	fmt.Printf("Nome do Aluno:\t%s\nSobrenome do Aluno:\t%s\nRG do Aluno:\t%s\nData de Admissão:\t%s\n", a.Nome, a.Sobrenome, a.RG, a.DataAdmissao)
}

func Exercicio1() {
	a := Alunos{"Jainy", "Martins Lemos", "MG1947571", time.Now().UTC()}
	a.detalhes()
}
