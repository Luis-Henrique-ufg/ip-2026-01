package main

import "fmt"

var menor int
var intermediario int
var maior int

func main() {

	fmt.Print("Digite a idade do menor: ")
	fmt.Scan(&menor)

	fmt.Print("Digite a idade do intermediário: ")
	fmt.Scan(&intermediario)

	fmt.Print("Digite a idade do maior: ")
	fmt.Scan(&maior)

	balinhasMenor := menor * 3
	balinhasIntermediario := intermediario * 2
	balinhasMaior := maior * 1

	totalBalinha := balinhasMenor + balinhasIntermediario + balinhasMaior

	fmt.Printf("Total de balinhas a comprar: %d\n", totalBalinha)
	fmt.Printf("%d balinhas para o neto de idade %d\n", balinhasMenor, menor)
	fmt.Printf("%d balinhas para o neto de idade %d\n", balinhasIntermediario, intermediario)
	fmt.Printf("%d balinhas para o neto de idade %d\n", balinhasMaior, maior)
}
