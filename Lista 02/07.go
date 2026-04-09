// Escreva um programa que leia três números inteiros distintos (assuma que o usuário irá digitar números distintos entre si)
// e os armazene nas variáveis A, B e C. Em seguida, descubra o menor valor e o armazene em uma variável denominada MENOR;
// o maior valor, coloque-o na variável MAIOR e o valor intermediário, na variável INTER.
// Imprima os valores em odem crescente, ou seja, imprima as variáveis MENOR, INTER e MAIOR.

package main

import "fmt"

var A, B, C int
var menor, maior, inter int

func main() {

	fmt.Println("Digite o primeiro número: ")
	fmt.Scan(&A)
	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&B)
	fmt.Println("Digite o terceiro número: ")
	fmt.Scan(&C)

	if A < B && A < C {
		menor = A
	} else if B < A && B < C {
		menor = B
	} else {
		menor = C
	}

	if A > B && A > C {
		maior = A
	} else if B > A && B > C {
		maior = B
	} else {
		maior = C
	}

	if A != menor && A != maior {
		inter = A
	} else if B != menor && B != maior {
		inter = B
	} else {
		inter = C
	}

	fmt.Printf("Em ordem crescente: %d %d %d\n", menor, inter, maior)
}
