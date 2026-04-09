package main

import "fmt"

var salarioMinimo float64
var EnergiaGasta float64

func main() {
	fmt.Println("Digite o valor do salário mínimo: ")
	fmt.Scan(&salarioMinimo)
	fmt.Println("Digite o valor da energia gasta: ")
	fmt.Scan(&EnergiaGasta)

	valorPorKw := (salarioMinimo * 0.7) / 100 //Valor de 1 KW = Salário mínimo * 0.7 / 100
	valorConsumo := EnergiaGasta * valorPorKw
	valorDesconto := valorConsumo * 0.90

	fmt.Printf("Custo por KW: R$ %.2f\n", valorPorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", valorConsumo)
	fmt.Printf("Custo do desconto: R$ %.2f\n", valorDesconto)

}
