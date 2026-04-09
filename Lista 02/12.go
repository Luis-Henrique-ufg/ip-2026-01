// Construa um programa que receba a idade de uma pessoa e classifique-a seguindo o critério apresentado a seguir.
// Considere que a idade é um valor inteiro e que será informada de forma válida.

package main

import f "fmt"

func main() {

	var idade int

	f.Println("Digite a sua idade: ")
	f.Scan(&idade)

	switch {
	case idade >= 0 && idade <= 2:
		f.Println("Recém-Nascido")

	case idade >= 3 && idade <= 11:
		f.Println("Criança")

	case idade >= 12 && idade <= 19:
		f.Println("Adolescente")

	case idade >= 20 && idade <= 55:
		f.Println("Adulto")

	case idade > 55:
		f.Println("Idoso")

	default:
		f.Println("Idade Inválida")
	}
}
