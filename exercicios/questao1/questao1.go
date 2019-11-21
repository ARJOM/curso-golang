package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Println("Informe três numeros:")
	fmt.Scan(&n1, &n2, &n3)
	media := (n1 + n2 + n3) / 3.0
	fmt.Printf("A média entre %d, %d e %d, é: %d \n", n1, n2, n3, media)
}
