programa 
{
  inclua biblioteca Matematica --> mat
  
  real x
  
  funcao inicio() 
  {
  escreva("Digite o salário a ser reajustado: ")  
  leia(x)

limpa()

  se(x <= 300.00)
  {
    x = 300.00 * 1.5
    mat.arredondar(x, 2)    
    escreva("Salário com reajuste: ", x, " reais\n")
  }
  senao
  {
    x = x * 1.3
    x = mat.arredondar(x, 2) 
    escreva("Salário com reajuste: ", x, " reais\n")
  }
  }
}
