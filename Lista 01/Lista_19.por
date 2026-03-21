programa 
{
  inclua biblioteca Matematica --> mat

  funcao inicio() 
  {
  inteiro n
  real cofre = 0.0 

  escreva("Digite um número inteiro maior que 1: ") 
  leia(n) 

  limpa()

  se(n <= 1)
  {
  escreva("Número inválido !\n")
  }
  senao 
  {
    para (inteiro k = 1; k <= n; k++)
    {
      cofre = cofre + (1.0 / k)
    }
    
    cofre = mat.arredondar(cofre, 6)
    escreva("O valor final do somatório é: ", cofre, "\n")
  }
  }
}
