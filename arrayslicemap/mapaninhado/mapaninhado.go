package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15143.00,
			"Gurgel Filho":   4523.21,
		},
		"J": {
			"Jośe Pereira": 4523.69,
		},
		"P": {
			"Paulo Filho": 8563.89,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
