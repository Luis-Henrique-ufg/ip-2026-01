// Escreva um programa que receba um número inteiro positivo de 3 casas e imprima o algarismo da casa das dezenas.
// Não se esqueça de testar para ver se o número informado tem realmente 3 casas.

package main

import "fmt"

func main() {

	var n int

	fmt.Println("Digite o seu número: ")
	fmt.Scan(&n)

	if n < 100 && n > 999 {
		fmt.Println("Número inválido")
	} else {
		d := (n / 10) % 10
		fmt.Printf("O algarismo da casa das dezenas do número %d é %d", n, d)
	}
}

// (200 / 10) % 10 = 20 % 10 = 0
// (371 / 10) % 10 = 37 % 10 = 7
