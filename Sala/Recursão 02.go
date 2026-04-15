package main

import "fmt"

func main() {
	numeros := []float64{4.5, 3.0, 2.5}
	resultado := somandoSlices(numeros)
	fmt.Printf("A soma do array é igual a %.2f", resultado)
}

func somandoSlices(numeros []float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	return numeros[0] + somandoSlices(numeros[1:])
}
