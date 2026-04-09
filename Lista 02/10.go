package main

import f "fmt"

func main() {

	var destino, retorno int

	f.Println("Escolha o se destino: |1. Norte| |2. Nordeste| |3. Centro-Oeste| |4. Sul|")
	f.Scan(&destino)
	f.Println("A passagem é Ida e Volta? |1. Sim| |2. Não|")
	f.Scan(&retorno)

	if (destino < 1 || destino > 4) || (retorno < 1 || retorno > 4) {
		f.Println("Valor Inválido")
	}

	var preco float64

	switch destino {
	case 1:
		if retorno == 1 {
			preco = 900
		} else {
			preco = 500
		}
	case 2:
		if retorno == 1 {
			preco = 650
		} else {
			preco = 350
		}
	case 3:
		if retorno == 1 {
			preco = 600
		} else {
			preco = 350
		}
	case 4:
		if retorno == 1 {
			preco = 550.00
		} else {
			preco = 300
		}
	}

	f.Printf("O preço da passagem é R$ %.2f", preco)
}
