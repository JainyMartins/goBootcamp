package main

import "fmt"

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

func main() {
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
