package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número inteiro (n>=3): ")
	fmt.Scan(&numero)

	if numero < 3 {
		fmt.Println("Número deve ser pelo menor 3")
	}

	termo1 := 0
	termo2 := 1

	if numero >= 1 {
		fmt.Print(termo1)
	}
	if numero >= 2 {
		fmt.Print("-", termo2)
	}

	for i := 3; i <= numero; i++ {
		proximo := termo1 + termo2
		fmt.Print("-", proximo)

		termo1 = termo2
		termo2 = proximo
	}

}
