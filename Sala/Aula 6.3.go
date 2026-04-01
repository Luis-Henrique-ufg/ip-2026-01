package main

import "fmt"

func main() {

	var valores [10]float64

	fmt.Println("Digite os dez valores reais: ")

	for i := 0; i < len(valores); i++ {
		fmt.Printf("Valor %d: ", i+1)
		fmt.Scan(&valores[i])
	}
	fmt.Println("Esses valore em ordem inversa são:\n ")
	for i := len(valores) - 1; i >= 0; i-- {
		fmt.Println(valores[i])
	}
}
