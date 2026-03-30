package main

import "fmt"

func main() {
	var x float64

	fmt.Println("Digite o salário a ser reajustado: ")
	fmt.Scan(&x)

	if x <= 300.00 {
		x = 300.00 * 1.5
		fmt.Printf("Salário com reajuste é = R$ %.2f", x)
	} else {
		x = x * 1.3
		fmt.Printf("Salário com reajuste é = R$ %.2f", x)
	}
}
