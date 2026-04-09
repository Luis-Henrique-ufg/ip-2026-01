package main

import (
	"fmt"
)

func main() {
	var n, quadrado int

	fmt.Println("Digite o seu número inteiro: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			quadrado = i * i
			fmt.Printf("%d ^ 2 = %d\n", i, quadrado)
		}
	}
}
