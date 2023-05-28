package main

import (
	"fmt"
	"time"
)

type Alunos struct {
	Nome         string
	Sobrenome    string
	RG           string
	DataAdmissao time.Time
}

func (a Alunos) detalhes() {
	fmt.Printf("Nome do Aluno:\t%s\nSobrenome do Aluno:\t%s\nRG do Aluno:\t%s\nData de Admissão:\t%s\n", a.Nome, a.Sobrenome, a.RG, a.DataAdmissao)
}

func main() {
	a := Alunos{"Jainy", "Martins Lemos", "MG1947571", time.Now().UTC()}
	a.detalhes()
}
