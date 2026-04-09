package main

import (
	"fmt"
	"math"
)

func main() {
	var figura int

	fmt.Println("Qual das três figuras você quer calcular a área e o volume?\nCone Reto(1)\nCilindro(2)\nEsfera(3)")
	fmt.Scan(&figura)

	var altura, raio, area, volume float64

	switch figura {
	case 1:
		fmt.Println("Raio: ")
		fmt.Scan(&raio)
		fmt.Println("Altura: ")
		fmt.Scan(&altura)

		volume = (math.Pi * (math.Pow(raio, 2)) * altura) / 3
		area = math.Pi * raio * math.Sqrt(math.Pow(raio, 2)+math.Pow(altura, 2))

	case 2:
		fmt.Println("Raio: ")
		fmt.Scan(&raio)
		fmt.Println("Altua: ")
		fmt.Scan(&altura)

		volume = math.Pi * math.Pow(raio, 2) * altura
		area = 2 * math.Pi * raio * altura

	case 3:
		fmt.Println("Raio: ")
		fmt.Scan(&raio)

		volume = 4.0 / 3.0 * math.Pi * math.Pow(raio, 3)
		area = 4 * math.Pi * math.Pow(raio, 2)

	default:
		fmt.Println("Figura inválida!")
	}
	fmt.Printf("\nÁrea = %.2f \n", volume)
	fmt.Printf("Volume = %.2f", area)
}
