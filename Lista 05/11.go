package main

import (
	"fmt"
	"math"
)

func main() {

	//var vetor = make([]float64, 4)
	var vetor = make([]float64, 100)

	Somatorio := 0.0
	for i := 0; i < len(vetor); i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < len(vetor)/2; i++ {

		indiceOposto := len(vetor) - 1 - i

		esquerda := vetor[i]
		direita := vetor[indiceOposto]

		//S=(b0-b99)^3 + (b1-b98)^3 + (b2-b97)^3 + . . . + (b49-b50)^3

		base := esquerda - direita
		resultado := math.Pow(base, 3)

		Somatorio += resultado

	}
	fmt.Printf("\nResultado final do somatório: %.2f", Somatorio)

}
