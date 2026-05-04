package main

import "fmt"

func main() {
	var vetor = make([]int, 100)

	for i := 0; i < len(vetor); i++ {
		vetor[i] = 2*i + 1
	}
	fmt.Println(vetor)
}
