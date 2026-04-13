package main

import "fmt"

func main() {
	const precoInicial = 6.00
	const ingressosIniciais = 130
	const despesas = 300.00
	const reducao = 0.60
	const acrescimo = 30

	var lucroMax float64
	var precoMax, ingressosMax float64

	fmt.Printf("%-10s %-15s %-15s\n", "Preço", "Ingressos", "Lucro")
	fmt.Println("==================================================")

	for i := 0; ; i++ {
		preco := precoInicial - float64(i)*reducao
		ingressos := ingressosIniciais + i*acrescimo
		receita := preco * float64(ingressos)
		lucro := receita - despesas

		if preco < 1.00 {
			break
		}

		fmt.Printf("R$ %-7.2f %-15d R$ %.2f\n", preco, ingressos, lucro)

		if i == 0 || lucro > lucroMax {
			lucroMax = lucro
			precoMax = preco
			ingressosMax = float64(ingressos)
		}
	}

	fmt.Println("\n--- Melhor cenário ---")
	fmt.Printf("Lucro máximo: R$ %.2f\n", lucroMax)
	fmt.Printf("Preço:        R$ %.2f\n", precoMax)
	fmt.Printf("Ingressos:    %.0f\n", ingressosMax)
}
