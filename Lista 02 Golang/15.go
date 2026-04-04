package main

import "fmt"

func main() {
	var dia, mes, ano int

	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	fmt.Print("Digite o mês: ")
	fmt.Scan(&mes)
	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	meses := [13]string{" ", "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro"}

	if mes < 1 || mes > 12 {
		fmt.Println("Mês inválido")
	} else if dia < 1 || dia > 31 {
		fmt.Println("Dia inválido")
	}

	fmt.Printf("%d de %s de %d", dia, meses[mes], ano) //Vá até a lista 'meses' e pegue o conteúdo que está dentro da caixinha de número 'mes'
}
