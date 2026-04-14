package main

import "fmt"

func buscaBinaria(list []int, alvo int) int {
	esquerdo := 0
	direito := len(list) - 1

	for esquerdo <= direito {
		meio := (direito + esquerdo) / 2
		elem_meio := list[meio]

		if elem_meio == alvo {
			return meio
		}
		if elem_meio > alvo {
			direito = meio - 1
		}
		if elem_meio < alvo {
			esquerdo = meio + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("Busca binaria...")
	list := []int{1, 5, 15, 20, 24, 45, 67, 76, 78, 100}
	posicao := buscaBinaria(list, 100)
	fmt.Println(posicao)
}
