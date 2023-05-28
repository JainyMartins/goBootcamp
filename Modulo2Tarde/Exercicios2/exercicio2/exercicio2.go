package exercicio2

import "fmt"

/*Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
produtos e devolver o valor do preço total.
As empresas têm 3 tipos de produtos:
- Pequeno, Médio e Grande.
Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto + 3% pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
adicional pelo envio de $2500.

Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos.
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
- Criar uma interface “Produto” que possua o método “CalcularCusto”
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
e preço, e devolva um Produto.
- Será necessário uma função “novaLoja” que retorne um Ecommerce.
- Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
custo adicional segundo o tipo de produto.

- Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
tenha)
- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
e adicioná-lo a lista da loja*/

type loja struct {
	produtos []produto
}

type produto struct {
	tipoProduto string
	nome        string
	preco       float64
}

func (p produto) CalcularCusto() float64 {
	switch {
	case p.tipoProduto == "pequeno":
		return p.preco
	case p.tipoProduto == "medio":
		return p.preco * 1.03
	case p.tipoProduto == "grande":
		return p.preco*1.06 + 2500.0
	default:
		return p.preco
	}
}

type Produto interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar()
}

func (l loja) Total(loja loja) float64 {
	soma := 0.0
	valorProduto := 0.0
	for _, valor := range loja.produtos {
		valorProduto = valor.CalcularCusto()
		soma += valorProduto
	}
	return soma
}

func (c *loja) Adicionar(produto produto) {
	c.produtos = append(c.produtos, produto)
}

func novoProduto(tipoProduto, nome string, preco float64) produto {
	return produto{tipoProduto: tipoProduto, nome: nome, preco: preco}
}

func novaLoja() loja {
	return loja{}
}

func Exercicio2() {
	//pequeno, médio e grande
	//pequeno - sem custo adicional, apenas do produto
	//medio - custo produto + 3% disponibilidade estoque
	//grande - custo produto + 6% disponibilidade estoque + custo add envio 2500

	loja := novaLoja()
	loja.Adicionar(novoProduto("pequeno", "morango", 9.9))
	loja.Adicionar(novoProduto("grande", "tv", 1050.9))
	fmt.Println(loja.Total(loja))
	fmt.Println(loja)
}
