//Escreva um programa que leia um número inteiro e diga se ele é ou não um número divisível por 5.

package main

import( 
    "fmt"
)
var x int

func main () {
    fmt.Println("Digite o número para verificação: ")
    fmt.Scan(&x)
    
    if x %5 == 0 || x %5 == 5 {
    fmt.Printf("O número %d é divisível por cinco", x)
    
    } else {
    fmt.Printf("O número %d não é divisível por cinco", x)
    }
}
        
    


