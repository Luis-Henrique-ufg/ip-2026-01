package main

import "fmt"

func main() {
	var preco float64
	var tipo, s int

	fmt.Println("Qual o preço normal do DVD? ")
	fmt.Scan(&preco)

	fmt.Println("Qual a categoria do DVD?\n1. Comum\n2. Lançamento")
	fmt.Scan(&tipo)

	fmt.Println("Hoje é que dia?\nDomingo(1)\nSegunda(2)\nTerça(3)\nQuarta(4)\nQuinta(5)\nSexta(6)\nSábado(7)")
	fmt.Scan(&s)

	if s == 2 || s == 3 || s == 5 {
		preco -= preco * 0.4
	}

	if tipo == 2 {
		preco += preco * 0.15
	}

	fmt.Printf("O preço final é R$ %.2f", preco)
}
