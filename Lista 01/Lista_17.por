programa 
{
  inclua biblioteca Matematica --> mat 

inteiro x, y, aux

  funcao inicio() 
  {
    escreva("Digite o primeiro número: ")
    leia(x)

    escreva("Digite o segundo número: ")
    leia(y)

    se (x % 2 != 0)
    {
      escreva("O primeiro número não é par")
    }
    senao
    {
      para(aux = 1; aux <= y; aux++)
      {
        escreva(x, " ")
        x = x + 2
      }
      escreva("\n")
    }

  }
}
