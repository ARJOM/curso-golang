package main

import (
	"fmt"
)

func main() {
	var palavra string
	var k int
	fmt.Println("Informe uma palavra e um número: ")
	fmt.Scan(&palavra, &k)
	letra := palavra[k : k+1]
	fmt.Print(letra)
}
