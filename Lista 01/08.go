// Um fabricante de latas deseja desenvolver um programa para calcular o custo de uma lata cilíndrica de alumínio, sabendo-se que o custo do alumínio por m2 é R$ 100,00.
package main

import (
	"fmt"
	"math"
)

func main() {
	var raio, altura, A_t, A_c, A_l float64

	fmt.Println("Digite o raio da lata em centímetros: ")
	fmt.Scan(&raio)

	fmt.Println("Digite a altura da lata em centímetros: ")
	fmt.Scan(&altura)

	A_c = 3.14159 * math.Pow(raio, 2)
	A_l = 2 * 3.14159 * raio * altura
	A_t = 2*A_c + A_l

	custo := A_t * 0.01 // Conversão

	fmt.Printf("O valor do custo é: R$ %.2f", custo)

}
