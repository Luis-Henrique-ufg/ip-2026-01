package main

import "fmt"

func main() {
	var primeiranota, segundanota float64
	var media float64
	var reprovado, exame, aprovado int
	var somaMedias, mediaClasse float64
	var totalalunos int

	for {
		fmt.Print("Digite a 1° nota ('123456' para encerrar):  ")
		fmt.Scan(&primeiranota)

		if primeiranota == 123456 {
			break
		}

		fmt.Print("Digite a 2° nota: ")
		fmt.Scan(&segundanota)

		totalalunos++

		media = (primeiranota + segundanota) / 2
		fmt.Printf("\nMédia: %.2f", media)

		if media <= 3 {
			fmt.Println("\nReprovado")
			reprovado++
		} else if media > 3 && media < 7 {
			fmt.Println("\nExame: ")
			exame++
		} else {
			fmt.Println("\nAprovado")
			aprovado++
		}

		somaMedias += media
		mediaClasse = somaMedias / float64(totalalunos)

	}

	if totalalunos == 0 {
		fmt.Println("\nNenhum aluno cadastrado.")
		return
	}

	fmt.Println("========== Resultados ==========")
	fmt.Printf("Total alunos: %d\n", totalalunos)
	fmt.Printf("Aprovados: %d\n", aprovado)
	fmt.Printf("Exame: %d\n", exame)
	fmt.Printf("Reprovados %d\n", reprovado)
	fmt.Printf("Média da classe: %.2f\n", mediaClasse)

}
