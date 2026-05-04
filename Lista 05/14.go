package main

import "fmt"

func main() {
	var vetor1 = make([]int, 10)
	var vetor2 = make([]int, 10)

	var resultante = make([]int, 20)

	fmt.Println("Digite 10 números inteiros: ")
	for i := 0; i < len(vetor1); i++ {
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("Digite mais 10 números inteiros: ")
	for j := 0; j < len(vetor2); j++ {
		fmt.Scan(&vetor2[j])
	}

	for i := 0; i < 10; i++ {
		resultante[i*2] = vetor1[i]
		resultante[i*2+1] = vetor2[i]
	}

	fmt.Println("\nVetor resultate intercalação: ")
	fmt.Println(resultante)

}
