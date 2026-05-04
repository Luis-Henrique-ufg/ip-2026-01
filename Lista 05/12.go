package main

import "fmt"

func main() {
	var notas = make([]int, 15)

	for i := 0; i < len(notas); i++ {
		fmt.Printf("Nota %d: ", i+1)
		fmt.Scan(&notas[i])
	}

	frequenciaAbs := map[int]int{}
	for _, v := range notas {
		frequenciaAbs[v]++
	}

	fmt.Printf("%-6s |%-15s |%-15s\n", "Notas", "Frequencia Abs", "Frquencia Relativa")

	totalAlunos := 15.0

	for notaAtual := 0; notaAtual <= 10; notaAtual++ {
		abs := frequenciaAbs[notaAtual]

		relativa := float64(abs) / totalAlunos

		fmt.Printf("%-6d | %-15d | %.2f\n", notaAtual, abs, relativa)
	}

}
