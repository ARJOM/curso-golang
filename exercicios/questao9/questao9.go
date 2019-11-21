package main

import "fmt"

func main() {
	const converter float64 = 3.6
	var km, h float64
	fmt.Println("Informe a distância percorrida em quilômetros: ")
	fmt.Scan(&km)
	fmt.Println("Informe a distância percorrida em quilômetros: ")
	fmt.Scan(&h)
	kmh := km / h
	ms := kmh / converter
	fmt.Printf("A viagem se deu em %.2f metros por segundo", ms)
}
