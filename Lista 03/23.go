package main

import "fmt"

func main() {
	var N int

	fmt.Print("Digite o número de termos N (N > 0):")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Erro: N deve ser um número positivo.")
		return
	}

	var soma float64 = 0.0

	for i := 1; i <= N; i++ {
		numerador := float64(1003 - 3*i)
		denomiador := float64(i)

		var termo float64

		if i%2 == 1 {
			termo = numerador / denomiador
		} else {
			termo = -numerador / denomiador
		}

		soma += termo
	}
	fmt.Printf("Soma dos %d primeirs termos: %.2f\n", N, soma)
}
