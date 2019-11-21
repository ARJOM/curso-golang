package main

import "fmt"

func main() {
	var n1, p1, n2, p2, n3, p3 float64
	fmt.Println("Informe o primeiro número e seu peso: ")
	fmt.Scan(&n1, &p1)
	fmt.Println("Informe o segundo número e seu peso: ")
	fmt.Scan(&n2, &p2)
	fmt.Println("Informe o terceiro número e seu peso: ")
	fmt.Scan(&n3, &p3)
	media := (n1*p1 + n2*p2 + n3*p3) / (p1 + p2 + p3)
	fmt.Printf("A média ponderada é %f", media)
}
