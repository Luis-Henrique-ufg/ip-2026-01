programa 
{
  inclua biblioteca Matematica --> mat 
  funcao inicio() {
    inteiro n 

    escreva("Digite o número: ")
    leia(n)

    limpa()

    se (n % 5 == 0 ou n % 3 == 0)
    {
      escreva("O número é divisível\n")
    }
    senao
    {
      escreva("O número não é divisível\n")
    }
  }
}
