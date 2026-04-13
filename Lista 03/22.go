package main

import "fmt"

func main() {
	soma := 0.0

	for i := 1; i <= 37; i++ {
		termo := float64((38-i)*(39-i)) / float64(i)
		soma += termo
	}

	fmt.Printf("S = %.2f\n", soma)
}
