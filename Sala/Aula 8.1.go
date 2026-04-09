package main

import "fmt"

func main() {

	var x, y float64

	fmt.Print("Digite os dois números: ")
	fmt.Scan(&x, &y)

	soma := soma(x, y)

	fmt.Println("A soma é = ", soma)
}

func soma(v1, v2 float64) float64 {
	return v1 + v2
}
