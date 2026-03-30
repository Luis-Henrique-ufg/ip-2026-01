package main

import "fmt"

func main() {

	var x, y float64

	fmt.Println("Digite o valor em graus Fahrenheit: ")
	fmt.Scan(&x)

	y = 5 * (y - 32) / 9

	fmt.Printf("%.2f °F é igual a %.2f °C", x, y)
}
