programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
  
  inteiro a, b, c  
  real delta

escreva("Qual o valor de a? ")
leia(a)

escreva("Qual o valor de b? ")
leia(b)

escreva("Qual o valor de c? ")
leia(c)

limpa()

delta = mat.potencia(b, 2) - 4 * a * c
delta = mat.arredondar(delta, 2)

escreva("O valor de Delta é: ", delta, "\n")

  }
}
