//Faça um programa que leia um número do tipo float e imprima sua raiz quadrada caso o mesmo seja positivo ou nulo. Caso ele seja negativo, mostre o seu quadrado.

package main

import( 
    "fmt"
    "math"
)
var x float64

func main () {
    
    fmt.Println("Digite o número: ")
    fmt.Scan(&x)
    
    if x >=0 {
        raiz := math.Sqrt(x)
        fmt.Printf("Resultado é: %2.f", raiz)
        
    } else if x < 0 { 
        x = x * x
        fmt.Printf("Resultado é %2.f", x)
    }
}
        
    


