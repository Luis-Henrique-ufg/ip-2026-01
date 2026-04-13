package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite o dividendo (N1): ")
	fmt.Scan(&n1)
	fmt.Print("Digite o divisor (N2, positivo): ")
	fmt.Scan(&n2)

	if n2 < 0 {
		fmt.Println("Erro, o divisor deve ser maior do que zero")
		return
	}
	if n1 < 0 {
		fmt.Println("Erro, o dividendo deve ser maior que zero")
	}
	quociente := 0
	resto := n1

	for resto >= n2 {
		resto -= n2
		quociente++
	}
	fmt.Printf("Quociente: %d, Resto: %d\n", quociente, resto)

}
