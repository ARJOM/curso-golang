package main

import "fmt"

func main() {
	var m, n int
	fmt.Println("Informe dois numeros")
	fmt.Scan(&m, &n)

	var cont int
	for i := m; i <= n; i++ {
		cont = 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				cont++
			}
		}
		if cont == 0 {
			fmt.Println(i)
		}
	}

}
