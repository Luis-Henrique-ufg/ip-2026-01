//Crie um programa para determinar se um número inteiro A é divisível por outro número B. Os valores devem ser fornecidos pelo usuário.

package main

import( 
    "fmt"
)

var x, y int

func main () {
    fmt.Println("Digite o primeiro número: ")
    fmt.Scan(&x)
    
    fmt.Println("Digite o segundo número: ")
    fmt.Scan(&y)
    
    if x % y == 0 {
        fmt.Printf("Os números %d e %d são divisíveis entre si", x, y)
    } else {
        fmt.Printf("Os números %d e %d NÃO são divisíveis ente si ", x, y)
    }
}
        
    


