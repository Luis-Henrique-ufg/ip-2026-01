//Crie um programa que leia um valor inteiro e diga se ele é positivo, negativo ou nulo.

package main

import "fmt"

var x int

func main () {
    
    fmt.Println("Digite o número para verificação: ")
    fmt.Scan(&x)
    
    if x < 0 {
        fmt.Printf("O número %d é NEGATIVO\n", x)
}   else if x > 0 {
        fmt.Printf("O número %d é POSITIVO\n", x)
}   else {
        fmt.Printf("O número %d é NULO", x)
    }
}
        
    


