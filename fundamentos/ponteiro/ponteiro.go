package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável i
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não possui aritimética de ponteiros
	// p++

	fmt.Println(p, &i, *p, i)
}
