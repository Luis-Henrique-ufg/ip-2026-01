package main

import "fmt"

func main() {
	var soma float64 = 0.0
	fatorial := int64(1) // 0!=1

	for k := 0; k < 20; k++ {
		numerador := float64(100 - k)
		termo := numerador / float64(fatorial)
		soma += termo

		fmt.Printf("%d \t%d\t\t%d\t\t%.2f\n", k, 100-k, fatorial, termo)

		fatorial *= int64(k + 1)
	}
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Soma dos 20 primeiros termos: %.2f\n", soma)
}
