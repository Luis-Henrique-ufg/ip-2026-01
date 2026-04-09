package main

import "fmt"

func main() {

	var matricula int
	var extras float64

	const salario = 788.00
	const horaextra = 10.00

	fmt.Println("Qual a matrícula do funcionário? ")
	fmt.Scan(&matricula)

	fmt.Println("Qual a quantidade de horas extras trabalhadas? ")
	fmt.Scan(&extras)

	salarioExtra := extras * horaextra
	bruto := 3*salario + salarioExtra

	descontos := 0.0

	if bruto > 1500 {
		descontos += bruto * 0.12
	}
	if bruto > 2000 {
		descontos += bruto * 0.20
	}
	liquido := bruto - descontos

	fmt.Printf("\nMatrícula: %d\n", matricula)
	fmt.Printf("Salário bruto: R$ %.2f\n", bruto)
	fmt.Printf("Descontos: R$ %.2f\n", descontos)
	fmt.Printf("Salário Líquido: R$ %.2f\n", liquido)
}
