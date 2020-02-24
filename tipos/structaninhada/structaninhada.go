package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) calculaTotal() float64 {
	soma := 0.0
	for _, item := range p.itens {
		soma += item.preco * float64(item.quantidade)
	}
	return soma
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 3, 16.0},
			item{3, 1, 4.0},
			item{4, 4, 30.02},
		},
	}

	fmt.Printf("O valor total do pedido é R$ %.2f", pedido.calculaTotal())

}
