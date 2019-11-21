package main

import (
	"fmt"
	"math"
)

func main() {
	var numero int
	fmt.Println("Informe um número:")
	fmt.Scan(&numero)
	dobro := numero * 2
	triplo := numero * 3
	quadruplo := numero * 4
	quadrado := math.Pow(float64(numero), 2.0)
	cubo := math.Pow(float64(numero), 3.0)
	raiz := math.Sqrt(float64(numero))
	fmt.Printf("Dobro: %d\nTriplo: %d\nQuadruplo: %d\nQuadrado: %.2f\nCubo: %.2f\nRaiz: %.2f", dobro, triplo, quadruplo, quadrado, cubo, raiz)
}
