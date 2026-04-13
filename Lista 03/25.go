package main

import (
	"fmt"
	"math"
)

// T_i = (-1)^i * \frac{2 ^ i }{( 15-i) ^2}

func main() {
	var S float64 = 0.0

	for i := 0; i <= 14; i++ {
		numerador := math.Pow(2, float64(i))     // 2 ^ i
		denomiador := math.Pow(float64(15-i), 2) // (15-i) ^ 2

		sinal := math.Pow(-1, float64(i)) // (-1) ^ i

		S += sinal * numerador / denomiador
	}
	fmt.Printf("O valor de S é: %.2f\n", S)
}
