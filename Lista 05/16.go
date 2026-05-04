package main

import "fmt"

func main() {

	vetor := []int{18, 18, 20, 21, 23, 19, 29, 45, 38, 42, 75, 19, 32, 76, 65, 87, 76, 54, 10, 20, 18, 18, 20, 21, 23, 19, 29, 45, 38, 42, 75, 19, 32, 76, 65, 87, 76, 54, 10, 20, 76, 10, 12, 85, 45, 46, 43, 10, 12, 40}

	var idade = make(map[int]int)

	for _, v := range vetor {
		idade[v]++
	}

	var moda int
	var maiorIdadeAtual int

	for key, value := range idade {
		if value > maiorIdadeAtual {
			moda = key
			maiorIdadeAtual = value
		}
	}

	fmt.Printf("Resultado oficial: a moda é %d anos (repetiu %d vezes)", moda, maiorIdadeAtual)

}
