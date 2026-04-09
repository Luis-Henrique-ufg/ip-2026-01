package main

import "fmt"

func main() {

	var A_1, r, n int
	var soma = 0

	fmt.Println("Qual o valor do primeiro termo? ")
	fmt.Scan(&A_1)

	fmt.Println("Qual o valor da razão? ")
	fmt.Scan(&r)

	fmt.Println("Qual o número de termos? ")
	fmt.Scan(&n)

	TermoAtual := A_1

	for i := 1; i <= n; i++ {
		soma += TermoAtual
		TermoAtual += r
	}
	fmt.Printf("A soma dos %d primeiros termos é %d\n", n, soma)

}
