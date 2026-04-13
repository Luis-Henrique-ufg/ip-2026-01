package main

import "fmt"

func main() {

	var H float64 = 0.0

	denominadores := 1
	numeradores := 1

	for i := 0; i <= 99; i++ {
		denominadores += 2
	}

	for i := 0; i <= 49; i++ {
		numeradores = 2*i + 1
		denominadores = i + 1

		H += float64(numeradores) / float64(denominadores)
	}

	fmt.Printf("H = %.2f", H)

}
