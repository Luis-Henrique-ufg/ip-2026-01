package main

import "fmt"

func main() {
	vetor := make([]int, 10)

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < len(vetor); i++ {
		_, err := fmt.Scan(&vetor[i])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
	}

	encontrou := false

	fmt.Println("\nResultados (números > 50):")
	for i := 0; i < len(vetor); i++ {
		if vetor[i] > 50 {
			fmt.Printf("Valor: %d na posição: %d\n", vetor[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Não existe nenhum número superior a 50.")
	}
}
