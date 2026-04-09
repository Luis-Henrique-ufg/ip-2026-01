package main

import "fmt"

func main() {
	var precoOriginal, percentualDesconto float64
	fmt.Print("Digite o valor original da mercadoria: R$ ")
	fmt.Scan(&precoOriginal)
	fmt.Print("Digite o percentual de desconto: ")
	fmt.Scan(&percentualDesconto)

	valorDesconto := precoOriginal * (percentualDesconto / 100)
	novoPreco := precoOriginal - valorDesconto

	fmt.Printf("Valor do desconto: R$ %.2f\n", valorDesconto)
	fmt.Printf("Novo valor a pagar: R$ %.2f\n", novoPreco)
}
