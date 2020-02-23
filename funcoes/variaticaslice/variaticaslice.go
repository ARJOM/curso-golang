package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Ricart", "Gabriella", "Francisco"}
	// Não funciona com arrays, por já terem tamanho fixo
	// aprovados := [...]string{"Ricart", "Gabriella", "Francisco"}
	imprimirAprovados(aprovados...)
}
