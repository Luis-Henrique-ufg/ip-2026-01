package main

import "fmt"

func main() {
	// fatorial de um número > 3! = 3 * 2 * 1 >  é a multiplicação pelo antecessor

	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero < 0 {
		fmt.Println("Erro, digite um número inteiro >= 0: ")
	}

	fatorial := 1

	for i := 1; i <= numero; i++ {
		fatorial *= i
	}

	fmt.Printf("O fatorial de %d é: %d", numero, fatorial)
}
