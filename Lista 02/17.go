package main

import "fmt"

func main() {

	var valor, consumo float64
	var tipo int

	fmt.Println("Selecione o tipo o sua categoria:\n1.Residencial\n2.Comercial\n3.Industrial")
	fmt.Scan(&tipo)
	fmt.Println("Digite o consumo total em metros cúbicos: ")
	fmt.Scan(&consumo)

	switch tipo {
	case 1:
		valor = (0.05 * consumo) + 5
	case 2:
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (consumo-80)*0.25
		}
	case 3:
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (consumo-100)*0.04
		}
	}
	fmt.Printf("Valor a pagar é R$: %.2f", valor)
}
