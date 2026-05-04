package main

import "fmt"

func main() {
	var vetor = make([]int, 4)

	var resultante = make([]int, 4)

	for i := 0; i < len(vetor); i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	for j := 0; j < len(vetor); j++ {
		if j%2 == 0 {
			resultante[j] = vetor[j] * 2
		} else {
			resultante[j] = vetor[j] * 3
		}
	}

	fmt.Println("O segundo vetor é: ")
	fmt.Println(resultante)

}
