package main

import "fmt"

var popular int
var geral int
var arquibancada int
var cadeiras int

func main() {
	fmt.Println("Digite o número de ingressos vendidos para a categoria popular: ")
	fmt.Scan(&popular)

	fmt.Println("Digite o número de ingressos vendidos para a categoria geral: ")
	fmt.Scan(&geral)

	fmt.Println("Digite o número de ingressos vendidos para a categoria arquibancada: ")
	fmt.Scan(&arquibancada)

	fmt.Println("Digite o número de ingressos vendidos para a categoria cadeiras:")
	fmt.Scan(&cadeiras)

	popular = popular * 1
	geral = geral * 5
	arquibancada = arquibancada * 10
	cadeiras = cadeiras * 20

	total := popular + geral + arquibancada + cadeiras
	fmt.Printf("A renda do jogo foi de: %d\n", total)
}
