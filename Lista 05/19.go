package main

import "fmt"

func main() {
	var vetor1 = make([]int, 10)
	var vetor2 = make([]int, 5)

	for i := 0; i < len(vetor1); i++ {
		fmt.Scan(&vetor1[i])
	}

	for j := 0; j < len(vetor2); j++ {
		fmt.Scan(&vetor2[j])
	}

	for k1, num1 := range vetor1 {
		achouDivisor := false

		for k2, num2 := range vetor2 {
			if num1%num2 == 0 {
				fmt.Printf("Divisível por %d na posição %d\n", num2, k2)
				achouDivisor = true
			}
		}

		if !achouDivisor {
			fmt.Println("Nenhum divisor encontrado no vetor 2")
		}
	}
}
