package main

import (
	"fmt"
	"math"
)

func main() {

	var altura, aresta float64

	fmt.Println("Digite a altura da pirâmide: ")
	fmt.Scan(&altura)

	fmt.Println("Digite o tamanho da aresta do hexágono que forma a base da pirâmide: ")
	fmt.Scan(&aresta)

	AreaBase := (3.0 * math.Pow(aresta, 2.0) * math.Sqrt(3.0)) / 2.0
	volume := (1.0 / 3.0) * AreaBase * altura

	fmt.Printf("O volume da pirâmide é: %2.f metros cúbicos", volume)

}
