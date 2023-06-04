package exercicio3

import "fmt"

/*Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

Precisamos de 3 estruturas:
- Produtos: nome, preço, quantidade.
- Serviços: nome, preço, minutos trabalhados.
- Manutenção: nome, preço.
Precisamos de 3 funções:
- Somar Produtos: recebe um array de produto e devolve o preço total (preço *
quantidade).
- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
tivesse trabalhado meia hora).
- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
tela (somando o total dos 3).*/

type Produtos struct{
	nome string
	preco float64
	quantidade int
}

type Servicos struct{
	nome string
	preco float64
	minutos int
}

type Manutencao struct{
	nome string
	preco float64
}

func somarProdutos(produtos []Produtos, c chan<- float64){
	total := 0.0
	for _, produto := range produtos {
		total += produto.preco * float64(produto.quantidade)
	}
	c <- total
}

func somarServicos(servicos []Servicos, c chan<- float64){
	total := 0.0
	var horasTrabalhadas float64
	for _, servico := range servicos{
		horasTrabalhadas = float64(servico.minutos)/60.0
		total += servico.preco * horasTrabalhadas
	}
	c <- total
}

func somarManutencoes(manutencoes []Manutencao, c chan<- float64){
	total := 0.0
	for _, manutencao := range manutencoes{
		total += manutencao.preco
	}
	c <- total
}

func Exercicio3(){
	produtos := []Produtos{
		{nome: "Café", preco: 10.99, quantidade: 2},
		{nome: "Pasta Dental", preco: 4.99, quantidade: 1},
	}
	servicos := []Servicos{
		{nome: "Serviço 1", preco: 100.99, minutos: 60},
		{nome: "Serviço 2", preco: 24.99, minutos: 120},
	}
	manutencoes := []Manutencao{
		{nome: "Manutenção 1", preco: 1000.00},
		{nome: "Manutenção 2", preco: 100.00},
	}

	totalProdutos := make(chan float64)
	totalServicos := make(chan float64)
	totalManutencao := make(chan float64)

	go somarProdutos(produtos, totalProdutos)
	go somarServicos(servicos, totalServicos)
	go somarManutencoes(manutencoes, totalManutencao)

	total := <-totalProdutos + <-totalServicos + <-totalManutencao
	fmt.Printf("Preço total: %.2f\n", total)
}