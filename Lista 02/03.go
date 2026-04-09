//Elabore um programa que leia dois números inteiros e efetue a adição dos mesmos. Caso o valor somado seja maior do que 20, o resultado a ser apresentado será a soma mencionada adicionada de adicionada de 8. Caso o valor somado seja menor ou igual a 20, deve-se subtrair 5 do mesmo para apresentá-lo em seguida.

package main

import "fmt"

var x, y int

func main () {
    
    fmt.Println("Digite o primeiro número: ")
    fmt.Scan(&x)
    
    fmt.Println("Digite o segundo número: ")
    fmt.Scan(&y)
    
    soma := x + y 
    
    if soma > 20 {
        soma = soma + 8
        fmt.Print("Resultado: ", soma)
    } else if soma <= 20 {
        soma = soma - 5
        fmt.Printf("Resultado: ", soma)
    }
}
        
    


