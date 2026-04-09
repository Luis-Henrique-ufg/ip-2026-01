package main

import (
	"fmt"
)

type pessoa struct {
	Nome      string
	Altura    float64
	PesoIdeal float64
}

func CalculoPesoIdeal(altura float64) float64 {
	return 72.7*altura - 58.0
}

func main() {

	var ppslice []pessoa

	for {
		var nome string
		var altura float64

		fmt.Print("Digite o nome (ou 'fim' para encerrar): ")
		fmt.Scan(&nome)

		if nome == "fim" || nome == "FIM" {
			break
		}

		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)

		person := pessoa{
			Nome:      nome,
			Altura:    altura,
			PesoIdeal: CalculoPesoIdeal(altura),
		}

		ppslice = append(ppslice, person)
	}
	fmt.Println("\n============== Lista de Pessoas ===============")
	fmt.Printf("%-20s %-15s %-10s\n", "Nome", "Altura", "Peso Ideal")
	fmt.Println("----------------------------------------------")

	for _, p := range ppslice {
		fmt.Printf("%-20s %-15.2f %-10.2f\n", p.Nome, p.Altura, p.PesoIdeal)
	}
	fmt.Println("----------------------------------------------")
	fmt.Printf("Total de pessoas cadastradas: %d\n", len(ppslice))
}
