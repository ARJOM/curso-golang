package main

import "fmt"

func main() {
	var numero int
	fmt.Println("Informe um número: ")
	fmt.Scan(&numero)
	antecessor := numero - 1
	sucessor := numero + 1
	fmt.Printf("Antecessor: %d\nSucessor: %d", antecessor, sucessor)
}
