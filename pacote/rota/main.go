package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	// É possível acessar o que está presente em rota.go
	// Isso se dá pois a visibilidade privada é a nível de pacote, e não de arquivo
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
