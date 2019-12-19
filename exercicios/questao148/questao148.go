package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(100)
}

func main() {
	n := numeroAleatorio() + 1

	fmt.Println(n)

	min := 0
	max := 100
	var palpite int
	fmt.Printf("Informe um número entre %d e %d", min, max)
	fmt.Scan(&palpite)

	for (max-min) != 2 && n != palpite {
		if palpite < min || palpite > max {
			continue
		} else if palpite < max && palpite > n {
			max = palpite
		} else if palpite > min && palpite < n {
			min = palpite
		}
		if (max - min) == 2 {
			continue
		}
		fmt.Printf("Escolha um numero entre %d e %d", min, max)
		fmt.Scan(&palpite)
	}

	if n == palpite {
		fmt.Println("Perdeu!!!")
	} else {
		fmt.Println("Ganhou!!!")
	}
}
