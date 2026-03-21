programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() 
  {
  inteiro a_1, r, n
  inteiro soma = 0
  inteiro termo_atual
  
  escreva("Qual o valor do primeiro termo? ")
  leia(a_1)

  escreva("Qual o valor da razão? ")
  leia(r)

  escreva("Qual o número de elementos? ")
  leia(n)

  limpa()

termo_atual = a_1

para (inteiro i = 1; i <= n; i++)
{
  soma = soma + termo_atual

  termo_atual = termo_atual + r
}

escreva("A soma dos ", n, " primeiros termos é: ", soma, "\n")
  }
}
