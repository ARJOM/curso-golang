package main

import "fmt"

func main() {
	var celsius float64
	fmt.Println("Informe um grau em Celsius")
	fmt.Scan(&celsius)
	kelvin := 273.15 + celsius
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("O grau informado representa %.2f graus Kelvin e %.2f graus Fahrenheit", kelvin, fahrenheit)

}
