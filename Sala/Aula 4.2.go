package main

import "fmt"

func main() {
	var qtdFarinha, valorFarinha float64
	var qtdFermento, valorFermento float64
	var qtdLuz, valorLuz float64
	var percentualImposto float64

	fmt.Print("Digite a quantidade de farinha (em kg):")
	fmt.Scan(&qtdFarinha)

	fmt.Print("Digite o valor da farinha (por kg): ")
	fmt.Scan(&valorFarinha)

	fmt.Print("Digite a quantidade de Fermento (em kg):")
	fmt.Scan(&qtdFermento)

	fmt.Print("Digite o valor do fermento (por kg):")
	fmt.Scan(&valorFermento)

	fmt.Print("Digite a quantidade de luz (em kWh):")
	fmt.Scan(&qtdLuz)

	fmt.Print("Digite o valor da luz (por kWh):")
	fmt.Scan(&valorLuz)

	fmt.Print("Digite o percentual de imposto: ")
	fmt.Scan(&percentualImposto)

	custoBase := (qtdFarinha * valorFarinha) + (qtdFermento * valorFermento) + (qtdLuz * valorLuz)
	valorImposto := custoBase * (percentualImposto / 100)
	precoCusto := custoBase + valorImposto
	precoVenda := precoCusto + (precoCusto * 0.30)

	fmt.Printf("Preço de custo R$ = %.2f\n", precoCusto)
	fmt.Printf("Preço de venda R$ = %.2f\n", precoVenda)
}
