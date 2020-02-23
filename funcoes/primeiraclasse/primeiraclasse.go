package main

import "fmt"

// Em Go as funções são consideradas cidadãs de primeira classe
// É possivel atribuir uma função a variáveis fora do main...
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(1, 5))

	// ... e também dentro do main
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 2))
}
