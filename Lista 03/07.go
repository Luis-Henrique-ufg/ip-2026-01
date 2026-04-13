package main

import "fmt"

func main() {
	var numeros float64
	var quantidade int
	var somaTotal float64
	var maior, menor float64
	var somaPares float64
	var qtdPares, qtdImpares int

	primeiro := true

	fmt.Println("Digite os números (ou digite 30.000 para finalizar): ")

	for {
		fmt.Scan(&numeros)

		if numeros == 30000 {
			break
		}
		quantidade++
		somaTotal += numeros

		if primeiro {
			maior = numeros
			menor = numeros
			primeiro = false
		} else if numeros > maior {
			maior = numeros
		}
		if numeros < menor {
			menor = numeros
		}
		if int(numeros)%2 == 0 {
			somaPares += numeros
			qtdPares++
		} else {
			qtdImpares++
		}
	}
	if quantidade == 0 {
		fmt.Println("Nenhum número válid foi digitado.")
		return
	}

	mediaTotal := somaTotal / float64(quantidade)

	var mediaPares float64

	if qtdPares > 0 {
		mediaPares = somaPares / float64(qtdPares)
	} else {
		mediaPares = 0
	}

	percentualImpares := float64(qtdImpares) / float64(quantidade) * 100

	fmt.Println("\n --- Resultados ---")
	fmt.Printf("a) Soma dos números digitados: %.2f\n", somaTotal)
	fmt.Printf("b) Quantidade de números digitados: %d\n", quantidade)
	fmt.Printf("c) Média dos números digitados: %2.f\n", mediaTotal)
	fmt.Printf("d) Maior número digitado: %.2f\n", maior)
	fmt.Printf("e) Menor número digitado: %2.f\n", menor)

	if qtdPares > 0 {
		fmt.Printf("f) Média dos números pares: %.2f\n", mediaPares)
	} else {
		fmt.Println("f) Não foram digitados números pares.")
	}
	fmt.Printf("g) Porcentagem de números ímpares: %.2f%%\n", percentualImpares)
}
