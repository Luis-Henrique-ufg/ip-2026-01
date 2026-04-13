package main

import "fmt"

func main() {
	const pi = 3.141592653589793
	fmt.Println("Raios (cm)      Volume (cm³)")
	fmt.Println("==============================")

	for i := 0; i <= 40; i++ {
		R := float64(i) * 0.5 // R varia de 0.0 a 20.0
		R3 := R * R * R
		volume := (4.0 / 3.0) * pi * R3
		fmt.Printf("%5.1f             %10.2f\n", R, volume)
	}

}

// V = \frac{4}{3} * pi * R^3
