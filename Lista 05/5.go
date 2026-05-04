package main

import "fmt"

func main() {
	var vetor = make([]int, 10)

	for i := 0; i < len(vetor); i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scan(&vetor[i])
	}

	menor := vetor[0]
	posicao := 0

	for j := 0; j < len(vetor); j++ {
		if vetor[j] < menor {
			menor = vetor[j]
			posicao = j
		}
	}
	fmt.Printf("\nO menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", menor, posicao)
}
