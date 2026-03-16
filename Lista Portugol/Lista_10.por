programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
    real a, b, c, d
    real delt_M
    
    escreva("Qual o valor do elemento a? ")
    leia(a)

    escreva("Qual o valor do elemento b? ")
    leia(b)

    escreva("Qual o valor do elemento c? ")
    leia(c)

    escreva("Qual o valor do elemento d? ")
    leia(d)

    delt_M = a * d - b * c 

    escreva("O valor do determinante é: ", delt_M,"\n")

  }
}
