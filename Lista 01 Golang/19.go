package main

import "fmt"

func main() {
	var n float64
	var cofre = 0.0

	fmt.Println("Digite um número interio maior que 1: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Número inválido")
	} else {
		for i := 1.0; i <= n; i++ {
			cofre = cofre + (1.0 / i)
		}
		fmt.Printf("O valor final do somatório é %.6f", cofre)
	}
}
