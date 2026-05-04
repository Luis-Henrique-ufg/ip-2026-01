package main

import "fmt"

func main() {
	var vetor = make([]int, 10)

	for i := 0; i < len(vetor); i++ {
		fmt.Scan(&vetor[i])
	}

	for k, v := range vetor {
		if ehPrimo(v) {
			fmt.Printf("O número %d - na posição %d do slice- é primo\n", v, k)
		} else {
			fmt.Printf("O número %d - na posição %d do slice- não é primo\n", v, k)
		}
	}

}

func ehPrimo(numero int) bool {
	if numero < 1 {
		return false
	}

	for divisor := 2; divisor < numero; divisor++ {
		if numero%divisor == 0 {
			return false
		}
	}
	return true
}
