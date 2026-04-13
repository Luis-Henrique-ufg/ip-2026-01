package main

import "fmt"

func main() {

	var n int
	for {
		fmt.Print("Digite um número inteiro: ")
		fmt.Scan(&n)

		if n <= 0 {
			fmt.Print("Programa Encerrado!")
			break
		}
		perfeito, raiz := QuadradoPerfeito(n)
		if perfeito {
			fmt.Printf("%d é quadrado perfeito! (raiz = %d)\n", n, raiz)
		} else {
			fmt.Printf("%d não é quadrado perfeito\n")
		}
	}

}

func QuadradoPerfeito(n int) (bool, int) {
	for i := 1; i <= n; i++ {
		if i*i == n {
			return true, i
		}
		if i*i > n {
			return false, 0
		}
	}
	return false, 0
}
