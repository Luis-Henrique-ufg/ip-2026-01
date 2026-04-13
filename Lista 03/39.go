package main

import "fmt"

func main() {
	var id, idMaisGordo, idMaisMagro int
	var peso, pesoMaisGordo, pesoMaisMagro float64

	fmt.Println("Boi 1: ")
	fmt.Print("ID:")
	fmt.Scan(&id)
	fmt.Print("Peso (kg): ")
	fmt.Scan(&peso)

	idMaisGordo, pesoMaisGordo = id, peso
	idMaisMagro, pesoMaisMagro = id, peso

	for i := 2; i <= 90; i++ {
		fmt.Printf("Boi %d\n", i)
		fmt.Print("ID: ")
		fmt.Scan(&id)
		fmt.Print("Peso (kg): ")
		fmt.Scan(&peso)

		if peso > pesoMaisGordo {
			idMaisGordo = id
			pesoMaisGordo = peso
		}

		if peso < pesoMaisMagro {
			idMaisMagro = id
			pesoMaisMagro = peso
		}
	}

	fmt.Printf("Boi mais gordo ID: %d | Peso: %.2f kg\n", idMaisGordo, pesoMaisGordo)
	fmt.Printf("Boi mais magro ID: %d | Peso: %.2f kg/n", idMaisMagro, pesoMaisMagro)
}
