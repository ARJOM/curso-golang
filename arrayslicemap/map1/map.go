package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[70520772482] = "Doravante"
	aprovados[11883576482] = "Marocas"
	aprovados[75354214796] = "Fulano"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[11883576482])
	delete(aprovados, 75354214796)
	fmt.Println(aprovados[75354214796])
}
