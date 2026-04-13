package main

import "fmt"

func main() {
	var totalGRaos uint64 = 0
	var graoNoQuadro uint64 = 1

	for quadrado := 1; quadrado <= 64; quadrado++ {
		totalGRaos += graoNoQuadro
		graoNoQuadro *= 2
	}
	fmt.Printf("Número toal de grãos de milho: %d\n", totalGRaos)

}
