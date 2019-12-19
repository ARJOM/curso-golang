package main

import "fmt"

func main() {
	var m, n int
	fmt.Println("Informe os numeros: ")
	fmt.Scan(&m, &n)
	ultimo := 1
	atual := 1
	proximo := ultimo + atual
	for atual < m {
		ultimo = atual
		atual = proximo
		proximo = atual + ultimo
	}
	for atual <= n {
		fmt.Println(atual)
		ultimo = atual
		atual = proximo
		proximo = atual + ultimo
	}
}
