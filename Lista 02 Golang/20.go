package main

import (
	"fmt"
)

func main() {
	var preco float64
	var pagamento int

	fmt.Println("Qual o preço normal de etiqueta? ")
	fmt.Scan(&preco)

	fmt.Println("\nEscolha a opção de pagamento:\n\n1. À vista: Dinheiro ou cheque\n2. À vista: Cartão de Crédito\n3. Em duas vezes\n4. Em três vezes\n")
	fmt.Scan(&pagamento)

	switch pagamento {
	case 1:
		preco -= preco * 0.1
	case 2:
		preco -= preco * 0.05
	case 3:
		preco = preco
	default:
		preco += preco * 0.1
	}
	fmt.Printf("\nO preço final = R$ %.2f", preco)
}
