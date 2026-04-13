package main

import "fmt"

func main() {
	fmt.Println("Índices de uma matriz 10x10:")

	for l := 0; l < 10; l++ {
		for c := 0; c < 10; c++ {
			fmt.Printf("(%d, %d)", l, c)
		}
		fmt.Println()
	}
}
