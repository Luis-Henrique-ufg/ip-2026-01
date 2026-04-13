package main

import "fmt"

func main() {

	var x, y int

	fmt.Print("Digite o 1° número: ")
	fmt.Scan(&x)
	fmt.Print("Digite o 2° número: ")
	fmt.Scan(&y)

	sinal := 1
	if (x < 0 && y > 0) || (x > 0 && y < 0) {
		sinal = -1
	}

	absX := x

	if absX < 0 {
		absX = -absX // 0 - absX
	}

	absY := y

	if absY < 0 {
		absY = -absY
	}

	menor := absX
	maior := absY

	if absX > absY {
		menor = absY
		maior = absX
	}

	resultado := 0

	for i := 0; i < menor; i++ {
		resultado += maior
	}
	if sinal == -1 {
		resultado = 0 - resultado
	}
	fmt.Printf("Rsultado da multiplicação: %d\n", resultado)

}
