package main

import (
	"fmt"
	//letra antes do pacote atribui à letra as funcionalidades do pacote
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é:", area)

	//bloco de constante
	const (
		a = 1
		b = 2
	)

	//bloco de variável
	var (
		c = 3
		d = 4
	)
	fmt.Println(c + d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "vish"

	fmt.Println(g, h, i)

}
