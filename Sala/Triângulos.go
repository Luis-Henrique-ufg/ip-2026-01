package main

import "fmt"

var a, b, c int

func main() {
	fmt.Println("Digite o primeiro lado do triângulo: ")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo lado do triângulo: ")
	fmt.Scan(&b)
	fmt.Println("Digite o terceiro lado do triângulo: ")
	fmt.Scan(&c)

	if a < b+c && b < a+c && c < a+b {
		if a == b && b == c {
			fmt.Println("O triângulo é equilátero.")
		} else if a == b || a == c || b == c {
			fmt.Println("O triângulo é isósceles.")
		} else {
			fmt.Println("O triângulo é escaleno.")
		}
	} else {
		fmt.Println("Não é um triângulo")
	}
}
