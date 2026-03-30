package main

import "fmt"

func main() {

	var a, b, c, d float64

	fmt.Println("Qual o valor do elemento a? ")
	fmt.Scan(&a)
	fmt.Println("Qual o valor do elemento b? ")
	fmt.Scan(&b)
	fmt.Println("Qual o valor do elemento c? ")
	fmt.Scan(&c)
	fmt.Println("Qual o valor do elemento d? ")
	fmt.Scan(&d)

	delta := a*d - b*c

	fmt.Printf("O valor do determinante é: %.2f\n", delta)
}
