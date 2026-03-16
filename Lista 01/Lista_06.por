programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() 
{
  real f, c
  caracter continuar
  
  faca
  {
    escreva("Digite a temperatura em graus Fahrenheit: ")  
  leia(f)
  f = mat.arredondar(f, 2)

  limpa()

  c = 5 * (f - 32) / 9
  c = mat.arredondar(c, 2)

  escreva(f, " Fahrenheit equivale a ", c, " graus Celsius\n")

  escreva ("Deseja converter outra temperatura? (S para sim/ N para não: ) ")
  leia(continuar)

  limpa()
  
  }
  enquanto (continuar == 'S' ou continuar == 's')

  escreva("Program encerreado. Até a próxima\n")
}
}