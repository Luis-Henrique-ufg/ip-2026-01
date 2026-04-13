package main

import "fmt"

func calcularDigito(digitos []int) int {
	soma := 0
	multiplicador := len(digitos) + 1

	for _, d := range digitos {
		soma += d * multiplicador
		multiplicador-- // decrementa
	}
	resto := soma % 11

	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func main() {
	cpf := make([]int, 9) //slice de nove dígitos

	fmt.Println("Digite os 9 primeiros dígitos do CPF: ")

	for i := 0; i < 9; i++ {
		fmt.Printf("Dígito %d: ", i+1)
		fmt.Scan(&cpf[i])
	}

	var d1, d2 int
	fmt.Print("Digite o 1° dígito verificador: ")
	fmt.Scan(&d1)
	fmt.Print("Digite o 2° dígito verificador: ")
	fmt.Scan(&d2)

	primeiro := calcularDigito(cpf)

	cpfCom10 := append(cpf, primeiro)
	segundo := calcularDigito(cpfCom10)

	fmt.Printf("\nDígito informados: %d%d\n", d1, d2)
	fmt.Printf("Dígitos calculados: %d%d\n", primeiro, segundo)

	if d1 == primeiro && d2 == segundo {
		fmt.Println("CPF Válido")

	} else {
		fmt.Println("CPF inválido")
	}
}
