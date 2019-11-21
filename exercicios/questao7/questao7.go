package main

import "fmt"

func main() {
	var base, altura float64
	fmt.Println("Informe a base e a altura: ")
	fmt.Scan(&base, &altura)
	area := base * altura / 2
	fmt.Printf("A área do triângulo é: %.2f", area)
}
