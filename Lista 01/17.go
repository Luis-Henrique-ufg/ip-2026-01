package main

import "fmt"

func main() {

	var x, y int

	fmt.Println("Digite o primeiro número: ")
	fmt.Scan(&x)

	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&y)

	if x%2 != 0 {
		fmt.Println("O primeiro número não é par!")
	} else {
		for i := 1; i <= y; i++ {
			fmt.Printf("%d ", x)
			x += 2
		}
	}
}
