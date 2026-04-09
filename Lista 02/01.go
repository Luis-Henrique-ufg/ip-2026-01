//Escreva um programa que leia um valor inteiro e diga se o número informado é par ou ímpar.

package main

import "fmt"

func main () {
    
    var x int
    
    fmt.Println("Digite um número: ")
    fmt.Scan(&x)
    
    if x % 2 == 0 {
        fmt.Printf("O número %d é PAR\n", x )
} else {
        fmt.Printf("O número %d é IMPAR\n", x)
    }
}


