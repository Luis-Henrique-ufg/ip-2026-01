//Faça um programa que leia três números inteiros e imprima tais números em ordem crescente.

package main

import "fmt"

func main() {

	var x, y, z int

	fmt.Print("Digite o 1° número: ")
	fmt.Scan(&x)
	fmt.Print("Digite o 2° número: ")
	fmt.Scan(&y)
	fmt.Print("Digite o 3° número: ")
	fmt.Scan(&z)

	if x > y {
		x, y = y, x
	}

	if x > z {
		x, z = z, x
	}

	if y > z {
		y, z = z, y
	}

	fmt.Printf("\nOs números em ordem crescente são: %d, %d, %d\n", x, y, z)
}
