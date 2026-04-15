package main

import "fmt"

func main() {
	var x, n int

	fmt.Print("Digite a base: ")
	fmt.Scan(&x)
	fmt.Print("Digite o expoente: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Digite um expoente positivo")
	}
	resultado := calculo(x, n)
	fmt.Printf("%d elevado a %d é igual a %d\n", x, n, resultado)

}

func calculo(x, n int) int {
	if n == 1 {
		return x
	}

	return x * calculo(x, n-1)
}
