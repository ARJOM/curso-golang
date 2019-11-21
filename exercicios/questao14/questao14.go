package main

import "fmt"

func main() {
	var palavra string
	fmt.Println("Digite uma palavra")
	fmt.Scan(&palavra)
	letras := len(palavra)
	fmt.Println(letras)
}
