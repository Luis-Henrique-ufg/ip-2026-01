package main

import (
	"fmt"
	"math"
)

func main() {

	var S float64 = 0.0

	for i := 0; i < 51; i++ {
		denominador := float64(2*i + 1)
		cubo := math.Pow(denominador, 3)

		termo := 1.0 / cubo

		if i%2 == 0 {
			S += termo
		} else {
			S -= termo
		}

	}

	produto := S * 32.0
	piAproximado := math.Cbrt(produto) //Cbrt = raíz cúbica
	fmt.Printf("Valor aproximado de pi com 51 termos: %.15f\n", piAproximado)
}
