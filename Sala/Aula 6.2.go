package main

import "fmt"

func main() {

	var valores [5]int

	soma := 0

	fmt.Println("Digite os cinco valores: ")

	for i := 0; i < len(valores); i++ {
		fmt.Printf("valor %d: ", i+1)
		fmt.Scan(&valores[i])
		soma += valores[i]
	}
	fmt.Printf("A soma desses valores é = %d", soma)
}
