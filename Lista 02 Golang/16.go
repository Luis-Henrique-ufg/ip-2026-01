// Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do segundo grau ( Ax^2 + Bx + C =0) e que calcule suas raízes.
// O programa deve mostrar, quando possível, o valor das raízes calculadas e a classificação das mesmas: “RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”

package main

import (
	"fmt"
	"math"
)

var A, B, C float64

func main() {

	fmt.Println("Digite o valor do coeficiente A: ")
	fmt.Scan(&A)
	fmt.Println("Digite o valor do coeficiente B: ")
	fmt.Scan(&B)
	fmt.Println("Digite o valor do coeficiente B: ")
	fmt.Scan(&C)

	if A < 0 {
		fmt.Println("O coeficiente A não pode ser zero")
	}

	delta := (math.Pow(B, 2)) - (4 * A * C)

	switch {
	case delta < 0:
		fmt.Println("Raíz Imaginárias")

	case delta == 0:
		fmt.Println("Raiz Única")

	default:
		x1 := (-B + math.Sqrt(delta)/2*A)
		x2 := (-B - math.Sqrt(delta)/2*A)
		fmt.Printf("X1 = %.2f\nX2 = %.2f", x1, x2)
	}
}
