// Faça um programa que indique se um número inteiro informado pelo usuário está compreendido entre
// 20 e 90 ou não. (20 e 90 não estão na faixa de valores).

package main

import "fmt"

var x int

func main() {
	fmt.Println("Digite o seu número: ")
	fmt.Scan(&x)

	if x > 20 && x < 90 {
		fmt.Printf("Sim, o seu número %d está entre 20 e 90 \n", x)
	} else {
		fmt.Printf("Não, o seu número %d não está entre 20 e 90 \n", x)
	}
}
