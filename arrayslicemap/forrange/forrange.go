package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numero := range numeros { // acessa o indice e o valor associado
		fmt.Printf("%d) %d\n", i, numero)
	}

	fmt.Println("")

	for i := range numeros { // pegando apenas o indice
		fmt.Println(i)
	}

	fmt.Println("")

	for _, num := range numeros { // ignorando o indice
		fmt.Println(num)
	}

}
