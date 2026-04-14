package main

import "fmt"

func buscaLinear(l []int, x int) int {

	for i := 0; i < len(l); i++ {

		if l[i] == x {

			return i
		}
	}
	return -1
}

func main() {
	minhaLista := []int{20, 5, 15, 24, 67, 45, 1, 76, 21, 11}
	numeroProcurado := 45

	indiceEncontrado := buscaLinear(minhaLista, numeroProcurado)

	if indiceEncontrado != -1 {
		fmt.Printf("O número %d foi encontrado no índice %d\n", numeroProcurado, indiceEncontrado)
	} else {
		fmt.Printf("O número %d não está na lista.\n", numeroProcurado)
	}
}
