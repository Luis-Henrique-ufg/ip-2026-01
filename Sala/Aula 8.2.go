package main

import "fmt"

func DescobrirMaior(a, b, c int) int {
	maior := a

	if b > maior {
		maior = b
	}

	if c > maior {
		maior = c
	}

	return maior
}

func main() {
	var num1, num2, num3 int

	fmt.Println("Digite três números inteiros:")
	fmt.Scan(&num1, &num2, &num3)

	resultado := DescobrirMaior(num1, num2, num3)

	fmt.Printf("\nO maior número entre os três é: %d\n", resultado)
}
