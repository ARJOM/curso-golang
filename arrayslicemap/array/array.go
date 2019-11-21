package main

import "fmt"

/*
Arrays são estruturas homogêneas, que possuem o mesmo tipo de dado armazenado neles.
Também são estruturas estáticas/fixas, ou seja, depois de definido seu tamanho ele não pode ser alterado
Além disso, Arrays são estruturas indexadas, que significa que cada elemento tem um índice (começando do 0)
*/

func main() {
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("A média é: %.2f\n", media)
}
