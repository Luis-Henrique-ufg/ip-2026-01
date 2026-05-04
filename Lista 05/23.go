package main

import "fmt"

func main() {

	janela := make([]int, 24)
	corredor := make([]int, 24)

	for {
		janelaCheia := estaCheio(janela)
		corredorCheio := estaCheio(corredor)

		if janelaCheia && corredorCheio {
			fmt.Println("\n[AVISO] O ônibus está complementamente cheio!")
			break
		}

		fmt.Println("\n[1] Vender Passagem")
		fmt.Println("[2] Encerrar Sistema")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		if opcao == 2 {
			fmt.Println("Sistema encerrado.")
			break
		} else if opcao != 1 {
			fmt.Println("[Erro] Opção inválida. ")
			continue
		}

		fmt.Println("\nQual a preferência do cliente?")
		fmt.Println("[1] Janela")
		fmt.Println("[2] Corredor")

		var preferencia int
		fmt.Scan(&preferencia)

		if preferencia == 1 {
			if janelaCheia {
				fmt.Println("[Aviso] Todas as poltronas na Janela estão ocupadas!")
			} else {
				fmt.Println("\nPoltronas Livres na Janela: ")
				mostrarLivres(janela)
				efetuarVenda(janela)
			}
		} else if preferencia == 2 {
			if corredorCheio {
				fmt.Println("\n[AVISO] Todas as poltronas no corredor estão oculpadas!")
			} else {
				fmt.Println("\nPoltronas livres no corredor: ")
				mostrarLivres(corredor)
				efetuarVenda(corredor)
			}
		} else {
			fmt.Println("[Erro] Preferência inválida.")
		}
	}
}

func estaCheio(setor []int) bool {
	for _, assento := range setor {
		if assento == 0 {
			return false
		}
	}
	return true
}

func mostrarLivres(setor []int) {
	fmt.Print("[ ")
	for i, assento := range setor {
		if assento == 0 {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println("]")
}

func efetuarVenda(setor []int) {
	fmt.Print("Digite o número da poltrona desejada: ")
	var numeroDigitado int
	fmt.Scan(&numeroDigitado)

	indice := numeroDigitado - 1

	if indice < 0 || indice >= len(setor) {
		fmt.Println("Erro. Número de poltronas inexistente.")

	} else if setor[indice] == 1 {
		fmt.Println("Erro. Essa poltrona já está ocupada! Venda cancelada.")
	} else {
		setor[indice] = 1
		fmt.Println("Sucesso. Passagem vendida com sucesso")
	}
}
