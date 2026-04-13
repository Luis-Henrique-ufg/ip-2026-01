package main

import (
	"fmt"
	"math"
)

func cossenoSerie(x float64, nTermos int) float64 {
	soma := 0.0

	for n := 0; n < nTermos; n++ {
	}

	soma = 0.0
	termo := 1.0 // termo para n=0
	soma += termo

	for n := 1; n < nTermos; n++ {
		termo *= -x * x / float64((2*n-1)*(2*n))
		soma += termo
	}
	return soma
}

func main() {
	var x float64
	fmt.Print("Digite o valor de x (em radianos): ")
	fmt.Scan(&x)

	cosSerie := cossenoSerie(x, 20)

	cosReal := math.Cos(x)

	diferenca := cosReal - cosSerie

	fmt.Printf("\n--- Resultados ---\n")
	fmt.Printf("Cosseno calculado pela série (20 termos): %.10f\n", cosSerie)
	fmt.Printf("Cosseno real (math.Cos): %.10f\n", cosReal)
	fmt.Printf("Diferença (real - série): %.10f\n", diferenca)
}
