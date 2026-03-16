programa 
{
inclua biblioteca Matematica --> mat 

  funcao inicio() 
  {

  real nota
  caracter conceito = ' '

  escreva("Nota: ")  
  leia(nota)

  se (nota >= 9.0 e nota <= 10.0)
  {
    conceito = 'A'
    escreva("NOTA = ", nota, " Conceito = ", conceito, "\n")
  }
  senao se(nota >= 7.5 e nota < 9.0)
  {
    conceito = 'B'
    escreva("NOTA = ", nota, " Conceito = ", conceito, "\n")
  }
  senao se (nota >= 6.0 e nota <7.5)
  {
    conceito = 'C'
    escreva("NOTA = ", nota, " Conceito = ", conceito, "\n")
  }
  senao se(nota >= 0.0 e nota < 6.0)
  {
    conceito = 'D'
    escreva("NOTA = ", nota, " Conceito = ", conceito, "\n")
  }
  }
}
