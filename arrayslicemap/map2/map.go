package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jacinto Medeiros": 15000.00,
		"Bezerra Brito":    12000.00,
		"Gomes Braga":      22000.00,
	}
	funcsESalarios["Mormai Son"] = 1350.0
	delete(funcsESalarios, "Ines Sistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
