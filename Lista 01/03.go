package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int

	// Lê os três números inteiros da entrada
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

	// Verifica se algum dos números tem mais de 1 dígito (ou seja, menor que 0 ou maior que 9)
	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
		return
	}

	// Faz a composição do número (n1 é centena, n2 é dezena, n3 é unidade)
	numero := (n1 * 100) + (n2 * 10) + n3

	// Calcula o quadrado do número
	quadrado := numero * numero

	// Imprime o resultado separado por vírgula e um espaço. 
	// O %d naturalmente não imprime zeros à esquerda.
	fmt.Printf("%d, %d\n", numero, quadrado)
}
