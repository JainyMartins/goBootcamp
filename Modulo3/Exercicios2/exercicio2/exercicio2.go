package exercicio2

import "fmt"

/*Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
mesmo endereço de memória no main do programa e nas funções.

Estruturas necessárias:
- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
- Produto: Nome, preço, quantidade.
Algumas funções necessárias:
- Novo produto: recebe nome e preço, e retorna um produto.
- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
o produto ao usuário.
- Deletar produtos: recebe um usuário, apaga os produtos do usuário.*/

type Usuario struct{
	nome string
	sobrenome string
	email string
	produtos []*Produto
}

type Produto struct{
	nome string
	preco float64
	quantidade int
}

func novoProduto (nome string, preco float64) *Produto{
	//Como vamos retornar um ponteiro, temos de retornar um endereço de memória para esta função, por isso usamos o &
	return &Produto{
		nome: nome,
		preco: preco,
	}
}

//A função abaixo espera um ponteiro para usuário como primeiro argumento. Portanto, precisamos passar um endereço de memória quando formos criar nosso usuário na main para chamar esta função
func adicionarProduto (usuario *Usuario, produto *Produto, quantidade int){
	produto.quantidade = quantidade
	usuario.produtos = append(usuario.produtos, produto)
}

func deletarProdutos (usuario *Usuario){
	usuario.produtos = nil
}

func Exercicio2(){
	//Criando um valor do tipo Usuario e obtendo o endereço de memória dele para atribuir ao usuário para que possamos usar adicionarprodutos, pois passamos um ponteiro, isso faz com que as alterações sejam refletidas no usuário original
	//Endereço de memória é o mesmo na main e nas funções
	usuario := &Usuario{
		nome: "João",
		sobrenome: "Martins",
		email: "joaomartins@gmail.com",
		produtos: []*Produto{},
	}

	fmt.Println("Usuário original:", usuario)
	produto1 := novoProduto("Calça", 39.99)
	adicionarProduto(usuario, produto1, 2)

	produto2 := novoProduto("Tênis", 99.9)
	adicionarProduto(usuario, produto2, 1)

	fmt.Println("Usuário com produtos adicionados:", usuario)

	deletarProdutos(usuario)

	fmt.Println("Usuário após deletarmos os produtos:", usuario)
}