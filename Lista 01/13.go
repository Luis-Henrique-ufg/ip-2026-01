package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Qual foi a nota (0 - 10)? ")
	fmt.Scan(&nota)

	if nota >= 9.0 && nota <= 10.0 {
		fmt.Printf("Nota: %.2f | Conceito A", nota)
	} else if nota >= 7.0 && nota < 9.0 {
		fmt.Printf("Nota: %.2f | Conceito B", nota)
	} else if nota >= 6.0 && nota < 7.5 {
		fmt.Printf("Nota: %.2f | Conceito C", nota)
	} else {
		fmt.Printf("Nota: %.2f | Conceito D", nota)
	}
}
