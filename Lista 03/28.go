package main

import (
	"fmt"
	"math"
)

func main() {
	// Passo 1: Inciar a variável que vai acumular a soma S
	var S float64 = 0.0
	// Passo 2: vamos usar um loop para somar 51 termos
	// O índice i vai de 0 até 50 (total 51 termos)
	for i := 0; i < 51; i++ {
		//O denominador é o (2 * i + 1) -ésimo número impar:
		denominador := float64(2*i + 1)
		// Calumos o cubo do denominador
		cubo := math.Pow(denominador, 3)
		// O termo atual é 1, sempre 1
		termo := 1.0 / cubo
		//Agora aplicamos o sinal: se i for par, sinal positivo
		// se i for impar, negativo.
		if i%2 == 0 {
			S += termo
		} else {
			S -= termo
		}

	}
	// Passo 3: após somar todos os 51 termos, multiplicamos S por 32
	produto := S * 32.0
	piAproximado := math.Cbrt(produto) //Cbrt = raíz cúbica
	fmt.Printf("Valor aproximado de pi com 51 termos: %.15f\n", piAproximado)
}
