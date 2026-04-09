package main

import "fmt"

func MEDIA(x, y, z float64) float64 {
	soma := x + y + z
	mediaFinal := soma / 3.0

	return mediaFinal
}

func main() {

	var valor1, valor2, valor3 float64
	fmt.Print("Digite a 1° nota: ")
	fmt.Scan(&valor1)
	fmt.Print("Digite a 2° nota: ")
	fmt.Scan(&valor2)
	fmt.Print("Digite a 3° nota: ")
	fmt.Scan(&valor3)

	resultado := MEDIA(valor1, valor2, valor3)

	// Imprimimos o resultado formatado com 2 casas decimais
	fmt.Printf("A média dos três valores é: %.2f\n", resultado)
}
