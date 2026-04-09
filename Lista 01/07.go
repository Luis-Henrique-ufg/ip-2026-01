package main

import "fmt"

func main() {

	var x, y, p, mm float64

	fmt.Println("Digite a temperatura em graus Fahrenheit: ")
	fmt.Scan(&x)

	fmt.Println("Digite a quantidade de chuva em polegadas: ")
	fmt.Scan(&p)

	y = 5 * (x - 32) / 9
	mm = p * 25.4

	fmt.Printf("%.2f °F é igual a %.2f °C\n", x, y)
	fmt.Printf("%.2f polegadas é igual a %2.f milímetros cúbicos de chuva", p, mm)
}
