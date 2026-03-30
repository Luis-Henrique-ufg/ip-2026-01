package main

import "fmt"

func main() {
	var codigo int
	var consumo float64
	var tipo string
	var valor float64

	fmt.Print("Você é um consumidor, residencial (R), comercial (C) ou industrial (I)? ")
	fmt.Scan(&tipo)

	fmt.Print("Qual foi o seu consumo em metros cubicos? ")
	fmt.Scan(&consumo)

	fmt.Print("Qual é o código indentificador da conta? ")
	fmt.Scan(&codigo)

	switch tipo {
	case "R", "r":
		valor = 5 + consumo*0.05
	case "C", "c":
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (consumo-80)*0.25 //Para quanto ele ultrapassou dos 80m^3
		}
	case "I", "i":
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (consumo-100)*0.04
		}
	}
	fmt.Printf("CONTA: %d\n", codigo)
	fmt.Printf("VALOR DA CONTA = R$ %.2f\n", valor)
}
