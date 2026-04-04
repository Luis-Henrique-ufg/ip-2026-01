package main

import "fmt"

func main() {

	var compra, lucro float64

	fmt.Println("Qual foi o valor da compra? ")
	fmt.Scan(&compra)

	_, err := fmt.Scan(&compra)
	if err != nil || compra < 0 {
		fmt.Println("Valor de compra inválido. Por favor, insira um valor numérico positivo.")
		return
	}

	if compra < 10.00 {
		lucro = 0.7
	} else if compra < 30.00 {
		lucro = 0.5
	} else if compra < 50.00 {
		lucro = 0.4
	} else {
		lucro = 0.3
	}
	precoVenda := compra * (1 + lucro)
	fmt.Printf("O valor da venda é R$ %.2f \n", precoVenda)
}
