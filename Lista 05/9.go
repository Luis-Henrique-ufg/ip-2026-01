package main

import "fmt"

func main() {
	var vetor = make([]float64, 10)

	var n float64

	fmt.Println("Digite a altura dos 10 atletas: ")

	soma := 0.0
	media := 0.0

	for i := 0; i < len(vetor); i++ {
		fmt.Printf("Altura %d: ", i+1)
		fmt.Scan(&n)

		vetor[i] = n
		soma += vetor[i]

	}

	media = soma / 10

	for k, v := range vetor {
		if v > media {
			fmt.Printf("\nA altura %.2f - no índice %d do vetor- está acima da média", v, k)
		}
	}
}

// Como eu posso somar +1 em uma variável a cada número 'n' digitado pelo usuário? eu não sei, pois se eu apenas coloco total++ eu somo 0 a 0. Posso fazer o total começar com o valor 1, talvez. Minha lógica não encontra respaldo em lugar algum.
// Eu preciso que cada n digitado seja == 1 para minha variável total.
// A minha variável média vai receber a soma de todos os n número digitados. Seria media += n?
// A condição é se a altura n digitada for maior que a média ela deve ser entrar no vetor: vetor[i] = n, se ela não for maior que a media ela não entra no vetor.
//
