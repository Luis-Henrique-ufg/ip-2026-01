package main

import "fmt"

func main() {
	var a, b, c, exercicios float64

	fmt.Println("Digite as três notas do aluno: ")
	fmt.Scan(&a, &b, &c)
	fmt.Println("\nQual foi a média dos exercícios desse aluno? ")
	fmt.Scan(&exercicios)

	aproveitamento := (a + (b * 2) + (c * 3) + exercicios) / 7

	if aproveitamento >= 9.0 {
		fmt.Printf("Aproveitamento %.2f |Conceito = A - Apovado!\n", aproveitamento)
	} else if aproveitamento >= 7.5 && aproveitamento < 9.0 {
		fmt.Printf("Aproveitamento %.2f |Coneito = B - Aprovado!\n", aproveitamento)
	} else if aproveitamento >= 6.0 && aproveitamento < 7.5 {
		fmt.Printf("Aproveitamento %.2f |Conceito = C - Aprovado!\n", aproveitamento)
	} else if aproveitamento >= 4.0 && aproveitamento < 6.0 {
		fmt.Printf("Aproveitamento %.2f |Conceito = D - Reprovado!\n", aproveitamento)
	} else if aproveitamento < 4.0 {
		fmt.Printf("Aproveitamento %.2f |Conceito = E - Reprovado!\n", aproveitamento)
	}
}
