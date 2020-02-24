package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Computador", 2700.00, []string{"Promoção", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id": 2, "nome": "Lápis", "preco": 1.5, "tags":["Escolar", "Escritório"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
