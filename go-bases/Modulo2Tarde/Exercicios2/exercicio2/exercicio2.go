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

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar(Produto)
}

type produto struct {
	tipoProduto string
	nome        string
	preco       float64
}

func (p produto) CalcularCusto() float64 {
	var custoAdicional float64
	switch {
	case p.tipoProduto == "pequeno":
		custoAdicional = 0.0
	case p.tipoProduto == "medio":
		custoAdicional = p.preco * 0.03
	case p.tipoProduto == "grande":
		custoAdicional = p.preco*0.06 + 2500.0
	default:
		custoAdicional = 0.0
	}
	return p.preco + custoAdicional
}

type loja struct {
	Produtos []Produto
}

func (l loja) Total() float64 {
	var total float64
	for _, produto := range l.Produtos {
		total += produto.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(produto Produto) {
	l.Produtos = append(l.Produtos, produto)
}

func novoProduto(tipoProduto, nome string, preco float64) Produto {
	return produto{
		tipoProduto: tipoProduto,
		nome:        nome,
		preco:       preco,
	}
}

func novaLoja() Ecommerce {
	return &loja{}
}

func Exercicio2() {
	//pequeno, médio e grande
	//pequeno - sem custo adicional, apenas do produto
	//medio - custo produto + 3% disponibilidade estoque
	//grande - custo produto + 6% disponibilidade estoque + custo add envio 2500

	loja := novaLoja()
	loja.Adicionar(novoProduto("pequeno", "morango", 9.9))
	loja.Adicionar(novoProduto("grande", "tv", 1050.9))
	fmt.Printf("Total da loja: R$%.2f\n", loja.Total())
	fmt.Println(loja)
}
