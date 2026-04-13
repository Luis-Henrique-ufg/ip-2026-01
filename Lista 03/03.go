package main

import "fmt"

func main() {

	var SCarlos float64
	var SJoao float64

	var meses = 0

	fmt.Println("Digite o salário de Carlos: ")
	fmt.Scan(&SCarlos)

	SJoao = 1.0 / 3.0 * SCarlos

	for SJoao < SCarlos {

		SCarlos = SCarlos * 1.02
		SJoao = SJoao * 1.05

		meses++
	}

	fmt.Printf("\nVao ser necessários %d meses para o salário de João igualar ao de Carlos", meses)

}
