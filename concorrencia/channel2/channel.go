package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)          // causa sincronismo
	go doisTresQuatroVezes(2, c) // causa assincronismo

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
	// fmt.Println(<-c)
}
