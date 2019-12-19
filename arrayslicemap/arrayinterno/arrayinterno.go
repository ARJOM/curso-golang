package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20) // Definindo um slice de 10 posições e array interno de vinte posições
	s2 := append(s1, 1, 2, 3) // Definindo um novo slice a partir do primeiro e adicionando valores ao fim
	fmt.Println(s1, s2)       // Exibindo os resultados

	s1[0] = 7           // Atualizando primeira posição do slice 1
	fmt.Println(s1, s2) // Mostrando que ambos os slices foram alterados, pois utilizam o mesmo array interno
}
