package main

import "fmt"

func main() {

	var idade, maiorque50, continuar, TotalPessoas int
	var altura, peso, idadesentre1020, maisdequarentakilos, somaAlturas1020 float64

	for {
		fmt.Print("Digite a idade: ")
		fmt.Scan(&idade)

		fmt.Print("Digite a altura em (m): ")
		fmt.Scan(&altura)

		fmt.Print("Digite o peso em (kg): ")
		fmt.Scan(&peso)

		TotalPessoas++

		if idade > 50 {
			maiorque50++
		}

		if idade > 10 && idade < 20 {
			somaAlturas1020 += altura
			idadesentre1020++
		}

		if peso < 40 {
			maisdequarentakilos++
		}
		fmt.Print("\nDeseja continuar? 1. Sim    2. Não   \n")
		fmt.Scan(&continuar)

		if continuar != 1 {
			break
		}
		fmt.Println()

	}

	fmt.Print("==== Resultados ===\n")
	fmt.Printf("Quantidade de pessoas acima dos 50 anos: %d\n", maiorque50)

	if idadesentre1020 > 0 {
		media1020 := somaAlturas1020 / float64(idadesentre1020)
		fmt.Printf("Media das alturas das pessoas entre 10 - 20 anos: %.2f\n", media1020)
	} else {
		fmt.Println("Média das alturas de pessoas entre 10 - 20 anos = 0")
	}

	if TotalPessoas > 0 {
		porcentagem := float64(TotalPessoas) * float64(maisdequarentakilos) / 100
		fmt.Printf("Porcentagem de pessoas com peso inferior a 40 quilos: %.2f%%\n", porcentagem)
	} else {
		fmt.Println("Porcentagem de pessoas com peso inferior a 40 quilos = 0.00%")
	}
}
