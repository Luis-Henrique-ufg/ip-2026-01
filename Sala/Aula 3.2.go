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
	if x > y && x > z {
		fmt.Printf("%d é o maior", x)
	} else if y > x && y > z {
		fmt.Printf("%d é o maior", y)
	} else {
		fmt.Printf("%d é o maior", z)
	}
}
