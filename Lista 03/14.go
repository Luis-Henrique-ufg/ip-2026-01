package main

import (
	"fmt"
	"math"
)

func main() {

	var n1, n2 int
	fmt.Print("Digite dois números: ")
	fmt.Scan(&n1, &n2)

	if n1 > n2 {
		n1, n2 = n2, n1
	}

	fmt.Printf("Números primos entre %d e %d:\n", n1, n2)

	for numero := n1; numero <= n2; numero++ {
		if Primo(numero) {
			fmt.Printf("%d ", numero)
		}
	}
}

func Primo(numero int) bool {
	if numero < 2 {
		return false
	}
	if numero == 2 {
		return true
	}
	if numero%2 == 0 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 3; i <= limite; i += 2 {
		if numero%i == 0 {
			return false
		}
	}
	return true
}
