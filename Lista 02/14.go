// Desenvolva um programa para calcular e imprimir o preço final de um carro. O valor do preço inicial de fábrica é fornecido pelo usuário. O carro pode ter as seguintes opções:
package main

import "fmt"

func main() {
	var preco float64
	var item float64

	fmt.Print("Digite o preço inicial de fábrica R$ ")
	fmt.Scan(&preco)

	total := preco

	fmt.Println("\nGostaria de alguns itens extra? (digite 0 para finalizar): ")
	fmt.Println("1. Ar condicionado")
	fmt.Println("2. Pintura Metálica")
	fmt.Println("3. Vidro Elétrico")
	fmt.Println("4. Direção Hidráulica")

	for {
		fmt.Print("\nDigite o número do item: ")
		fmt.Scan(&item)

		if item == 0 {
			break
		}

		switch item {
		case 1:
			total += 1750.00
			fmt.Println("Ar condicionado adicionado!")
		case 2:
			total += 800.00
			fmt.Println("Pintura metálica adicionada!")
		case 3:
			total += 1200.00
			fmt.Println("Vidro elético adicionado!")
		case 4:
			total += 2000.00
			fmt.Println("Direção hidráulica adiconada")
		default:
			fmt.Println("Escolha inválida, tente novamente.")
		}
	}
	fmt.Printf("\nO preço final do carro é R$ %.2f\n", total)
}
