package main

import "fmt"

func main() {

	var horas, minutos, segundos int

	fmt.Println("Digite a quantidade de horas: ")
	fmt.Scan(&horas)
	fmt.Println("Digite a quantidade minutos: ")
	fmt.Scan(&minutos)
	fmt.Println("Digite a quantidade de segundos: ")
	fmt.Scan(&segundos)

	TotalSegundos := (horas * 60 * 60) + (minutos * 60) + segundos

	fmt.Printf("O tempo em segundos é = %d s", TotalSegundos)
}
