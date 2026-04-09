// Escreva um programa que receba o valor de x, calcule e imprima o valor de f(x), dado que: f(x) = 8 / (2-x)

package main

import f "fmt"

func main() {

	var x float64

	f.Println("Digite o valor de x: ")
	f.Scan(&x)

	if x == 2 {
		f.Println("A divisão por zero não é possível")
	} else {

		fx := 8 / (2 - x)
		f.Printf("O valor de f(x) é = %.2f", fx)
	}
}
