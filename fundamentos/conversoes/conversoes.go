package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// É obrigatório fazer o cast para o tipo esperado
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//fmt.Println("Teste " + 123) Não funciona, por ser de tipos diferentes

	/*
		Diferente de outras linguagens, fazer o cast de um numero para string vai resultar em seu correspondente
		na tabela ASCII do numero em questão
	*/

	fmt.Println("Teste " + string(123)) //Teste {

	// O jeito correto de converter um inteiro para string
	fmt.Println("Teste " + strconv.Itoa(123)) // Teste 123

	// O jeito correto de converter uma string para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(122-num, reflect.TypeOf(num))

	//
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
