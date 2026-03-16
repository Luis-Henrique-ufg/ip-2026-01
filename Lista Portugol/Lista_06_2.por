programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() 
{
  real f, c
  inteiro i
  
  para (i = 1; i <= 5; i++)
  {
    escreva("Digite a temperatura em graus Fahrenheit: ")  
  leia(f)
  f = mat.arredondar(f, 2)

  limpa()

  c = 5 * (f - 32) / 9
  c = mat.arredondar(c, 2)

  escreva(f, " graus Fahrenheit equivale a ", c, " graus Celsius\n")

  }
}
}