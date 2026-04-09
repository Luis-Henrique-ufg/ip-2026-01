package main

import "fmt"

func main() {

	var x int

	fmt.Println("Esse programa verifica se o seu número é divisível por 5 ou por 3")
	fmt.Println("=================")
	fmt.Println("Digite o número: ")
	fmt.Scan(&x)

	if x%5 == 0 || x%3 == 0 {
		fmt.Printf("O número %d é divisível", x)
	} else {
		fmt.Printf("O número %d não é divisível", x)
	}
}
