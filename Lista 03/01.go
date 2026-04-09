package main

import "fmt"

func main() {
	var base, expoente int

	fmt.Print("Digite a base: ")
	fmt.Scan(&base)

	fmt.Print("Digite o expoente: ")
	fmt.Scan(&expoente)

	resultado := 1

	for i := 1; i <= expoente; i++ {
		resultado = resultado * base
	}

	fmt.Printf("Resultado: %d", resultado)

}
