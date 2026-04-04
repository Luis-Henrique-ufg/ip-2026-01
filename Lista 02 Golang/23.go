package main

import "fmt"

func main() {

	var idade int

	fmt.Println("Digite a idade: ")
	fmt.Scan(&idade)

	switch {
	case idade < 16:
		fmt.Println("Não Eleitor")
	case idade >= 18 && idade < 65:
		fmt.Println("Eleitor Obrigatório")
	case idade > 16 && idade < 18 || idade >= 65:
		fmt.Println("Eleitor Facultativo")
	default:
		fmt.Println("Idade Inválida!")
	}
}
