package main

import (
	"fmt"
	"sort"
)

type Empregado struct {
	ID    int
	Meses int
}

func main() {

	var firma = make([]Empregado, 0, 100)

	for {
		var idDigitado, mesesDigitado int
		fmt.Scan(&idDigitado, &mesesDigitado)

		if idDigitado == 0 && mesesDigitado == 0 {
			break
		}

		if len(firma) >= 100 {
			break
		}

		novoCracha := Empregado{ID: idDigitado, Meses: mesesDigitado}
		firma = append(firma, novoCracha)
	}

	sort.Slice(firma, func(i, j int) bool {
		return firma[i].Meses < firma[j].Meses
	})

	limite := 3

	if len(firma) < 3 {
		limite = len(firma)
	}

	for i := 0; i < limite; i++ {
		fmt.Printf("%d° lugar -> Empregado ID: %d | Tempo: %d meses\n", i+1, firma[i].ID, firma[i].Meses)
	}
}
