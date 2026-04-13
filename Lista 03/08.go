package main

import "fmt"

func main() {

	numeros := make([]int, 20)

	var Total int

	for i := 0; i < 20; i++ {
		numeros[i] = i + 1
		Total += numeros[i]

	}
	fmt.Printf("Números: %d", numeros)
	fmt.Printf("\nSoma dos números: %d", Total)
}
