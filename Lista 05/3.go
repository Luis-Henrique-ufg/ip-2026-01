package main

import "fmt"

func main() {
	arr := make([]int, 10)

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	var pares []int
	var impares []int
	somaPares := 0

	for _, num := range arr {
		if num%2 == 0 {
			pares = append(pares, num)
			somaPares += num
		} else {
			impares = append(impares, num)
		}
	}

	fmt.Println("\n--- Resultados ---")
	fmt.Println("a) Números pares digitados:", pares)
	fmt.Println("b) Soma dos números pares digitados:", somaPares)
	fmt.Println("c) Números ímpares digitados:", impares)
	fmt.Println("d) Quantidade de números ímpares digitados:", len(impares))
}
