package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14
	var raio float64
	fmt.Println("Informe o raio da circunferência: ")
	fmt.Scan(&raio)
	area := PI * math.Pow(raio, 2.0)
	comprimento := 2 * PI * raio
	fmt.Printf("A área da circunferência é: %.2f\nO comprimento da circunferência é: %.2f", area, comprimento)
}
