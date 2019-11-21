package main

import "fmt"

func main() {
	var kb float64
	const conv float64 = 1024
	fmt.Println("Informe uma quantidade de kilobytes: ")
	fmt.Scan(&kb)
	byt := kb * conv
	bit := byt * conv
	mega := kb / conv
	giga := mega / conv
	fmt.Printf("Kb: %f\nBit: %f\nByte: %f\nMegabyte: %f\nGigabyte: %f", kb, bit, byt, mega, giga)

}
