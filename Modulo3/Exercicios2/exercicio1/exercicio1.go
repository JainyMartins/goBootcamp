package exercicio1

import "fmt"

/*Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
para as funções:
- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
senha
E devem implementar as seguintes funcionalidades:
- mudar o nome: me permite mudar o nome e sobrenome
- mudar a idade: me permite mudar a idade
- mudar e-mail: me permite mudar o e-mail
- mudar a senha: me permite mudar a senha*/

type ISetUsuario interface {
	mudarNomeESobrenome(*Usuario)
	mudarIdade(*Usuario)
	mudarEmail(*Usuario)
	mudarSenha(*Usuario)
}

type Usuario struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

func (u *Usuario) mudarNomeESobrenome(nome, sobrenome string) {
	u.nome = nome
	u.sobrenome = sobrenome
}

func (u *Usuario) mudarIdade(idade int) {
	u.idade = idade
}

func (u *Usuario) mudarEmail(email string) {
	u.email = email
}

func (u *Usuario) mudarSenha(senha string) {
	u.senha = senha
}

func Exercicio1() {
	var p *Usuario
	Usuario1 := Usuario{"Jainy", "Martins Lemos", 25, "jainyestefany@hotmail.com", "12345"}
	p = &Usuario1
	fmt.Println("Endereço de Memória do Ponteiro: ", &p)
	fmt.Println("Usuário: ", Usuario1)
	p.mudarEmail("jainy@gmail.com")
	p.mudarIdade(27)
	p.mudarNomeESobrenome("Jaina", "Lemos")
	p.mudarSenha("58767")
	fmt.Println("Endereço de Memória do Ponteiro após Alterações: ", &p)
}
