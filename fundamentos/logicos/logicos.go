package main

import "fmt"

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTV50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2 // ou exclusivo
	comprarSorvete := trabalho1 || trabalho2

	return comprarTV50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}
