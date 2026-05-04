package main

import "fmt"

func main() {
	vetor1 := make([]int, 10)
	vetor2 := make([]int, 5)

	fmt.Println("Digite os 10 números do primeiro vetor:")
	for i := 0; i < len(vetor1); i++ {
		fmt.Printf("Vetor 1 [%d]: ", i)
		fmt.Scan(&vetor1[i])

	}

	fmt.Println("\nDigite os 5 números do segundo vetor:")
	somaVetor2 := 0
	for i := 0; i < len(vetor2); i++ {
		fmt.Printf("Vetor 2 [%d]: ", i)
		fmt.Scan(&vetor2[i])

		somaVetor2 += vetor2[i]
	}

	var resultante1 []int
	var resultante2 []int

	for _, num := range vetor1 {
		if num%2 == 0 {
			resultante1 = append(resultante1, num+somaVetor2)
		} else {
			resultante2 = append(resultante2, num+somaVetor2)
		}
	}

	fmt.Println("\n--- Resultados ---")
	fmt.Println("Primeiro vetor resultante (Pares + Soma V2):", resultante1)
	fmt.Println("Segundo vetor resultante (Ímpares + Soma V2):", resultante2)
}
