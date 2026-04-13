package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	quadrado := 0

	for i := 1; i <= n; i++ {
		quadrado = i * i
		fmt.Print(" ", quadrado)
	}

}
