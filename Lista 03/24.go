package main

import (
	"fmt"
	"math"
)

func fatorial(n int) float64 {
	resultado := 1.0
	for i := 2; i <= n; i++ {
		resultado *= float64(i)
	}
	return resultado
}

func senMacLaurin(a float64) float64 {
	sen := 0.0
	sinal := 1.0

	for exp := 1; exp <= 7; exp += 2 {
		termo := math.Pow(a, float64(exp)) / fatorial(exp)
		sen += sinal * termo
		sinal *= -1
	}
	return sen
}

func main() {
	fmt.Printf("%-10s %-15s %-15s\n", "Ângulo", "Mac-Laurin", "math.Sin")
	fmt.Println("------------------------------------------")

	for a := 0.0; a <= 6.3; a += 0.1 {
		senCalculado := senMacLaurin(a)
		senReal := math.Sin(a)
		fmt.Printf("%-10.1f %-15.6f %-15.6f\n", a, senCalculado, senReal)
	}
}
