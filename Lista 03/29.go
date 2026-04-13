package main

import "fmt"

func main() {
	var n int
	fmt.Print("Até que número o somatório deve ir? ")
	fmt.Scan(&n)

	var soma int
	for i := 1; i <= n; i++ {
		soma += i
	}
	fmt.Printf("Soma de 1 a %d = %d\n", n, soma)
}
