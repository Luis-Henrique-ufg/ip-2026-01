//Um comerciante calcula o valor da venda, tendo em vista a tabela a seguir:
//Valor da Compra
//Valor < R$ 10,00
//R$ 10,00 <= Valor < R$ 30,00
//R$ 30,00 <= Valor < R$ 50,00
//Valor >= 50,00

//Valor da Venda
//Lucro de 70%
//Lucro de 50%
//Lucro de 40%
//Lucro de 30%

//Escreva um programa que leia o valor da compra e imprima o valor da venda. Cuidado com valor inválido de compra!

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
