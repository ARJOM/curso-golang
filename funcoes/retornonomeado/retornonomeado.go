package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	a, b := trocar(2, 3)
	fmt.Println(a, b)

	// pode usar os nomes usados para retorno, já que foram definidos no escopo da função
	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
