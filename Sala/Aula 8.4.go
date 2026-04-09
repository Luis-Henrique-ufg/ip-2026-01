package main

import "fmt"

func Fatorial(n int) int {

	if n <= 0 {
		return 1
	}

	resultado := 1

	for i := n; i > 0; i-- {
		resultado = resultado * i
	}

	return resultado
}

func main() {
	var numero int

	fmt.Print("Digite um número para calcular o fatorial: ")
	fmt.Scan(&numero)

	resposta := Fatorial(numero)

	fmt.Printf("O fatorial de %d! é: %d\n", numero, resposta)
}
