package main

import (
	"fmt"
	"math"
)

func converterBase10(n int) int {
	resultado := 0
	posição := 0

	for n > 0 {
		digito := n % 10
		resultado += digito * int(math.Pow(8, float64(posição)))
		posição++
		n /= 10
	}
	return resultado
}

func validarOctal(n int) bool {
	for n > 0 {
		digito := n % 10
		if digito >= 8 {
			return false
		}
		n /= 10
	}
	return true
}

func main() {
	var n int
	fmt.Print("Digite um número na base 8: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Número inválido! Digite um número positivo.")
		return
	}

	if !validarOctal(n) {
		fmt.Println("Número inválido! Em base 8 os dígitos vão de 0 a 7.")
		return
	}
	fmt.Printf("%d em base 8 = %d em base 10\n", n, converterBase10(n))
}
