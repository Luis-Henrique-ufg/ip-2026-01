package main

import "fmt"

func main() {

	var notas [3]float64

	soma := 0.0

	fmt.Println("Digite as três notas")

	for i := 0; i < len(notas); i++ {
		fmt.Printf("Nota %d:", i+1)
		fmt.Scan(&notas[i])
		soma += notas[i]
	}

	media := soma / float64(len(notas)) //Precisou converter len(notas) porque "media" é float e len(notas) era inteiro

	fmt.Printf("A média das notas é: |%.2f|\n", media)

	fmt.Println("Essas notas estão acima da média: ")

	for i := 0; i < len(notas); i++ {
		if notas[i] > media {
			fmt.Printf("|%.2f|", notas[i])
		}
	}
}
