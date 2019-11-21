package main

import "fmt"

func main() {
	var real, cotacao float64
	fmt.Println("Informe o valor em reais a ser convertido, e a cotação atual do dólar: ")
	fmt.Scan(&real, &cotacao)
	// dolar := real / cotacao
	dolar := real * cotacao
	fmt.Printf("O valor informado equivale a %.2f dolares", dolar)
}
