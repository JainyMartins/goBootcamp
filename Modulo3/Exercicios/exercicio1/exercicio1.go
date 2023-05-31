package exercicio1

/*Uma empresa que vende produtos de limpeza precisa de:
1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
de produtos comprados, separados por ponto e vírgula (csv).
2. Deve possuir o ID do produto, preço e a quantidade.
3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.*/

import (
	"fmt"
	"os"
)

type IProduto interface {
	MontarString() string
}

type IListaProdutos interface {
	Adicionar(IProduto)
	MontarConteudo() string
}

type Produto struct {
	id         int
	preco      float64
	quantidade int
}

func (p Produto) MontarString() string {
	s1 := fmt.Sprintf("%d", p.id)
	s2 := fmt.Sprintf("%.2f", p.preco)
	s3 := fmt.Sprintf("%d", p.quantidade)

	stringProduto := string(s1) + ";" + string(s2) + ";" + string(s3) + "\n"
	return stringProduto
}

type ListaProdutos struct {
	Produtos []IProduto
}

func (l *ListaProdutos) Adicionar(produto IProduto) {
	l.Produtos = append(l.Produtos, produto)
}

func (l ListaProdutos) MontarConteudo() string {
	var stringFinal string
	for i, produto := range l.Produtos {
		if i == 0 {
			stringFinal += "id;" + "preço;" + "quantidade" + "\n" + produto.MontarString()
		} else {
			stringFinal += produto.MontarString()
		}
	}
	return stringFinal
}

func novoProduto(id int, preco float64, quantidade int) IProduto {
	return Produto{
		id:         id,
		preco:      preco,
		quantidade: quantidade,
	}
}

func novaListaProdutos() IListaProdutos {
	return &ListaProdutos{}
}

func Exercicio1() {
	arquivo := novaListaProdutos()
	arquivo.Adicionar(novoProduto(1, 2.4, 5))
	arquivo.Adicionar(novoProduto(2, 5.5, 7))
	arquivo.Adicionar(novoProduto(3, 7.9, 3))
	conteudoString := arquivo.MontarConteudo()

	//Salvar arquivo em csv
	conteudo := []byte(conteudoString)
	erro := os.WriteFile("./listaProdutos.csv", conteudo, 0644)
	if erro != nil {
		fmt.Println("Erro ao escrever o arquivo", erro)
	}
}
