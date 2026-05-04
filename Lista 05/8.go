package main

import (
	"fmt"
	"math"
)

func main() {
	var vetor = make([]float64, 15)

	var n int

	fmt.Println("Digite os 15 números: ")

	for i := 0; i < len(vetor); i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&n)
		if n < 0 {
			vetor[i] = -1
		} else {
			vetor[i] = math.Sqrt(float64(n))
		}
	}

	fmt.Println("--- As raízes dos números inteiros positivos digitados são: ---")
	fmt.Printf("\n%.2f", vetor)

}
