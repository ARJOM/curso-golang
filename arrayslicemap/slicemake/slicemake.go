package main

import "fmt"

func main() {
	s := make([]int, 10) // Criando slice de um array de 10 posições
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // Criando slice de 10 posições de array interno de 20 posições
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // Adicionando elementos ao slice até o limite do array interno
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // Adicionando elementos além do limite do array interno
	fmt.Println(s, len(s), cap(s))
}
