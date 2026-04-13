package main

import "fmt"

func main() {
	var b, n int

	fmt.Print("Digite o a base: ")
	fmt.Scan(&b)
	fmt.Print("Digite o expoente: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Digite um valor maior que um para o expoente")
		return
	}

	if b < 2 {
		fmt.Println("Digite uma base maior que 2")
		return
	}
	resultado := 1
	for i := 1; i <= n; i++ {
		resultado *= b

	}

	fmt.Printf("%d elevado a %d = %d\n", b, n, resultado)

}
