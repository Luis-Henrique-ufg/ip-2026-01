package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Erro: o número deve ser positivo.")
		return
	}

	if num == 0 {
		fmt.Println("Binário: 0")
		return
	}

	binario := ""
	temp := num

	for temp > 0 {
		resto := temp % 2
		if resto == 1 {
			binario = "1" + binario
		} else {
			binario = "0" + binario
		}
		temp = temp / 2
	}
	fmt.Printf("O número %d em binário é %s\n", num, binario)

}

// 10/2 == 5, resto == 0
// 5/2 == 2, resto == 1
// 2/2 == 1, resto == 0
// 1/2 == 0, resto == 1
//2° iteração: temp=5, resto=1 → binario = "1" + "0" = "10"; temp=2
//1° iteração: temp=10, resto=0 → binario = "0" + "" = "0"; temp=5
//3° iteração: temp=2, resto=0 → binario = "0" + "10" = "010"; temp=1
//4° iteração: temp=1, resto=1 → binario = "1" + "010" = "1010"; temp=0
