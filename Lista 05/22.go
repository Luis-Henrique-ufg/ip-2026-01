package main

import "fmt"

func main() {
	codigos := make([]int, 10)
	saldos := make([]float64, 10)

	for i := 0; i < 10; i++ {
		for {
			fmt.Print("Digite o código da conta: ")
			var codigoDigitado int
			fmt.Scan(&codigoDigitado)

			if encontrarConta(codigos, codigoDigitado) != -1 {
				fmt.Println("Erro. Código já cadastrado.")
			} else {
				codigos[i] = codigoDigitado
				break
			}
		}

		fmt.Print("Digite o saldo inicial: R$ ")
		fmt.Scan(&saldos[i])
	}

	for {
		fmt.Println("1. Efetuar depósito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar o ativo bancário")
		fmt.Println("4. Finalizar o programa")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		if opcao == 1 {
			fmt.Print("\nDigite o código da conta para depósito: ")
			var codBusca int
			fmt.Scan(&codBusca)

			indice := encontrarConta(codigos, codBusca)

			if indice == -1 {
				fmt.Println("Erro. Conta não encontrada.")
			} else {
				fmt.Print("Digite o valor do depósito: R$ ")
				var valor float64
				fmt.Scan(&valor)

				saldos[indice] += valor
				fmt.Printf("Sucesso. Depósito realizado! Novo saldo: R$ %.2f\n",
					saldos[indice])
			}
		} else if opcao == 2 {
			fmt.Print("\nDigite o código da conta para saque: ")
			var codBusca int
			fmt.Scan(&codBusca)

			indice := encontrarConta(codigos, codBusca)

			if indice == -1 {
				fmt.Println("Erro. Conta não encontrada.")
			} else {
				fmt.Print("Digite o valor do saque: R$ ")
				var valor float64
				fmt.Scan(&valor)

				if saldos[indice] >= valor {
					saldos[indice] -= valor
					fmt.Printf("Sucesso. Saque realizado! Novo saldo: R$ %2f\n", saldos[indice])
				} else {
					fmt.Println("Erro. Saldo insuficiente.")
				}
			}
		} else if opcao == 3 {
			somaTotal := 0.0
			for _, saldo := range saldos {
				somaTotal += saldo
			}
			fmt.Printf("\nAtivo Bancário. O banco possui um total de R$ %2f\n", somaTotal)

		} else if opcao == 4 {
			fmt.Println("Encerrando o sistema...")
			break
		} else {
			fmt.Println("\nErro. Opção inválida.")
		}
	}
}

func encontrarConta(listaCodigos []int, codigoBuscado int) int {
	for i, codigo := range listaCodigos {
		if codigo == codigoBuscado {
			return i
		}
	}
	return -1
}
