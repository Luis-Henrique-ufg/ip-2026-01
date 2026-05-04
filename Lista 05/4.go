package main

import "fmt"

func main() {

	var vetor = make([]int, 10)

	for i := 0; i < len(vetor); i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	contagem := map[int]int{}
	for _, v := range vetor {
		contagem[v]++
	}

	fmt.Println("--- Elementos Repetidos ---")

	repete := false
	for k, v := range contagem {
		if v > 1 {
			fmt.Printf("\nO elemento %d se repetiu %d vezes", k, v)
			repete = true
		}
		if !repete {
			return
		}
	}
}
