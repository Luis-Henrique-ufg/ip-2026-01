package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Println("Digite o valor de A: ")
	fmt.Scan(&A)

	fmt.Println("Digite o valor de B: ")
	fmt.Scan(&B)

	fmt.Println("Digite o valor de C: ")
	fmt.Scan(&C)

	delta := math.Pow(B, 2) * A * C

	fmt.Printf("O valor de Delta é: %2.f\n", delta)

}
