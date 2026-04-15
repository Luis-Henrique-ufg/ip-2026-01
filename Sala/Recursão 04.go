// Um problema típico em ciência da computação consiste em converter um número da sua forma decimal para a forma binária.
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro postivo: ")
	fmt.Scan(&n)

	num := n
	fmt.Print("Binário: ")
	binario(num)

}

func binario(n int) {
	if n == 0 {
		return
	}
	binario(n / 2)
	fmt.Print(n % 2)
}
