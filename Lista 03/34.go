package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite o 1° número: ")
	fmt.Scan(&n1)
	fmt.Print("Digite o 2° número: ")
	fmt.Scan(&n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Println("Erro, digite números positivos")
		return
	}

	a, b := n1, n2
	for b != 0 {
		resto := a % b
		a = b
		b = resto
	}
	mdc := a
	mmc := n1 * n2 / mdc
	fmt.Printf("O MMC de %d e %d é %d\n", n1, n2, mmc)
}

// 14 / 4 = 3, resto == 2----  a == 4; b == 2
// 4 / 2 = 2, resto == 0 ----  a == 2; b == 0
// 2 / 0 ------ a == mdc
// mmc = a * b / mdc ---- 14 * 2 / 2 == 28
