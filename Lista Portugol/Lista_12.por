programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
  
  inteiro h, s, pacotes, total

  escreva("Qual foi o tempo de uso em horas? ")
  leia(h)

limpa()

  pacotes = h / 3
  s = h % 3
  total = (pacotes * 10) + (s * 5)

  escreva("O valor a pagar é: ", total, "\n")

  }
}
