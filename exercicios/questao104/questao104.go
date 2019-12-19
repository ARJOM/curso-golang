package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe um número")
	fmt.Scan(&n)
	var soma int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			soma += i
		}
	}
	if soma == n {
		fmt.Println("É perfeito")
	} else {
		fmt.Println("Não é perfeito")
	}
}
