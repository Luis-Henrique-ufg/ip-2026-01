package main

import (
	"fmt"
	"sort"
)

func main() {
	var vetor = make([]int, 10)

	for i := 0; i < len(vetor); i++ {
		fmt.Scan(&vetor[i])

		predacoPreenchido := vetor[:i+1]

		sort.Ints(predacoPreenchido)
	}

	fmt.Println("Vetor final ordenado: ", vetor)
}
