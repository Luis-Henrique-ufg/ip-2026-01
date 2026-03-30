package main

import "fmt"

func main() {

	var horas int

	fmt.Println("Quai foi o tempo de uso da charrete em horas? ")
	fmt.Scan(&horas)

	pacote := horas / 3
	resto := horas % 3

	total := float64((pacote * 10.0) + (resto * 5.0))

	fmt.Printf("O valor a pagar é R$ %.2f", total)
}
