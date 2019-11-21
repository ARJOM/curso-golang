package main

import (
	"fmt"
	"math"
)

func main() {
	var cateto1, cateto2 float64
	fmt.Println("Informe o valor dos catetos: ")
	fmt.Scan(&cateto1, &cateto2)
	hipotenusa := math.Pow(cateto1, 2.0) + math.Pow(cateto2, 2.0)
	hipotenusa = math.Sqrt(hipotenusa)
	fmt.Printf("O valor da hipotenusa é: %f", hipotenusa)
}
