package main

import "fmt"

func converterBase16(n int) string {
	digitos := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	if n == 0 {
		return "0"
	}

	resultado := ""

	for n > 0 {
		resto := n % 16
		resultado = digitos[resto] + resultado
		n = n / 16
	}
	return resultado
}

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Número inválido!")
		return
	}

	fmt.Printf("%d em base 16 = %s\n", n, converterBase16(n))
}
