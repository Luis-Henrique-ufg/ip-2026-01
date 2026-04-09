package main

import "fmt"

func main() {

	var x, y float64
	var pesox, pesoy float64

	fmt.Print("Digite a 1° nota e seu peso: ")
	fmt.Scan(&x, &pesox)
	fmt.Print("Digite a 2° nota: ")
	fmt.Scan(&y, &pesoy)

	soma := pesox + pesoy

	media := (x*pesox + y*pesoy) / soma

	fmt.Printf("Média Ponderada: %.2f", media)
}
