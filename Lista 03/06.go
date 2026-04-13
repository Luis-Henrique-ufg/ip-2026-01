package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&numero)

	if numero <= 0 {
		fmt.Println("Por favor, digite um número inteiro positivo.")
		return
	}

	encontrado := false
	var a, b, c int

	for i := 1; i*(i+1)*(i+2) <= numero; i++ {
		if i*(i+1)*(i+2) == numero {
			encontrado = true
			a, b, c = i, i+1, i+2
			break
		}
	}

	if encontrado {
		fmt.Printf("%d é um número triangular!\n", numero)
		fmt.Printf("Porque: %d * %d * %d = %d\n", a, b, c, numero)
	} else {
		fmt.Printf("%d NÃO é um número triangular.\n", numero)
	}
}
