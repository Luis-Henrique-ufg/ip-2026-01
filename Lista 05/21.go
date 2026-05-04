package main

import "fmt"

func main() {
	var vetor = make([]float64, 10)
	var codigo int

	for i := 0; i < len(vetor); i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Println("[0] Finalizar Programa")
	fmt.Println("[1] Mostrar oredem direta")
	fmt.Println("[2] Mostrar ordem inversa")
	fmt.Scan(&codigo)

	if codigo == 0 {
		fmt.Println("Programa encerrado pelo usuário;")
		return
	} else if codigo == 1 {
		for i := 0; i < len(vetor); i++ {
			fmt.Printf("Posição %d = %.2f\n", i, vetor[i])
		}
	} else if codigo == 2 {
		for i := len(vetor) - 1; i >= 0; i-- {
			fmt.Printf("Posição %d = %.2f\n", i, vetor[i])
		}
	} else {
		fmt.Println("\nCódigo inválido! Digite apenas 0, 1 ou 2")
	}
}
