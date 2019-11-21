package main

import "fmt"

func main() {
	var lado float64
	fmt.Println("Informe o lado do quadrado: ")
	fmt.Scan(&lado)
	perimetro := lado * 4
	area := lado * lado
	fmt.Printf("O perimetro do quadrado informado é: %v m\nA área do quadrado informado é: %v m²", perimetro, area)
}
