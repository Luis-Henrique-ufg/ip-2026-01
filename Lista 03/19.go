package main

import "fmt"

func main() {
	fmt.Println("Índices acima da diagonal principal de uma matriz 10x10 são: ")

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			fmt.Printf("(%d %d)", i, j)
		}
		fmt.Println()
	}
}
