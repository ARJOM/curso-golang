package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define uma parte de um array.
	s2 := a2[1:3] // slice com elementode do indice 1 ao 3, sem acessar o indice 3
	fmt.Println(s2)
	a2[1] = 6 // slice é um ponteiro
	fmt.Println(s2)

	s3 := a2[:2] // novo slice, mas apontando para o mesmo array
	fmt.Println(a2, s3)

	s4 := s2[:1] // slice de um slice
	fmt.Println(s4)

}
