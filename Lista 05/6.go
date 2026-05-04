package main

import "fmt"

func main() {

	var vetor = make([]int, 100)

	for i := 0; i < len(vetor); i++ {
		vetor[i] = 100 - i
	}

	fmt.Println(vetor)
}
