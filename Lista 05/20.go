package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var jogadas = make([]int, 20)

	var frequencia = make(map[int]int)

	for i := 0; i < len(jogadas); i++ {
		faceSorteadas := rand.Intn(6) + 1

		jogadas[i] = faceSorteadas
	}

	fmt.Printf("\nResultado dos 20 lançamentos:\n%v\n", jogadas)

	for _, v := range jogadas {
		frequencia[v]++
	}

	fmt.Println("Frequência de cada face")

	for face := 1; face <= 6; face++ {
		qtd := frequencia[face]

		fmt.Printf("A face %d foi sorteada %d vezes.\n", face, qtd)
	}
}
