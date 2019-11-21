package main

import "fmt"

func main() {
	var polegadas float64
	const conv float64 = 2.54
	fmt.Println("Informe um numero em polegadas: ")
	fmt.Scan(&polegadas)
	centimetros := conv * polegadas
	fmt.Printf("Em centímetros: %.2f", centimetros)
}
