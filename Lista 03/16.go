package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	fmt.Print("Qual o tamanho da sequência: ")
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("Erro. Digite um número maior do que 3.")
		return
	}

	fmt.Printf("Série de Fetuccine com %d números\n", n)
	fmt.Printf("%d %d ", a, b)

	anterior2 := a
	anterior1 := b

	for i := 3; i < n; i++ {
		var atual int
		if i%2 == 1 {
			atual = anterior1 + anterior2
		} else {
			atual = anterior1 - anterior2
		}

		fmt.Printf(" %d", atual)

		anterior2 = anterior1
		anterior1 = atual
	}
}
