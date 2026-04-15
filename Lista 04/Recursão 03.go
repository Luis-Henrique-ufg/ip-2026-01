// Dado um array de inteiros e o seu número de elementos, inverta a posição dos seus elementos.
package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4}
	invertendo(numeros, 0, len(numeros)-1)
	fmt.Println(numeros)
}
func invertendo(numeros []int, inicio int, fim int) {
	if inicio >= fim {
		return
	}
	numeros[inicio], numeros[fim] = numeros[fim], numeros[inicio]

	invertendo(numeros, inicio+1, fim-1)

}
