package main

import "fmt"

func main() {
	var vetor = make([]int, 50)

	vetor[0] = 1
	vetor[1] = 1

	fmt.Println("--- Aqui estão os 50 primeiros termos da série de Fibonacci --- ")
	for i := 2; i < len(vetor); i++ {
		vetor[i] = vetor[i-2] + vetor[i-1]
	}

	fmt.Println(vetor)
}
