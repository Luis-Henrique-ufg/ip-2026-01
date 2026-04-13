package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x)

	S := 0.0
	fatorial := 1.0

	for i := 0; i <= 20; i++ {
		termo := x / fatorial

		if i%2 == 0 {
			S += termo
		} else {
			S -= termo
		}
		fatorial *= float64(i + 1)
	}
	fmt.Printf("O valor do somatório S é: %.2f\n", S)

}
