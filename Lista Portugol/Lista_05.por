programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {

    inteiro conta
    real consumo
    caracter tipo
    real valor

    escreva("Você é um consumidor, residencial (R), comercial (C) ou industrial (I)? ")
    leia(tipo)

    escreva("Qual foi o seu consumo em metros cubicos? ")
    leia(consumo)

    escreva("Qual é o código indentificador da conta? ")
    leia(conta)

    limpa()

    se (tipo == 'R')
    {
      valor = 5 + consumo * 0.05
    }

    se (tipo == 'C')
    {
      se (consumo <= 80)
      {
        valor = 500
      }
      senao
      {
        valor = 500 + (consumo - 80) * 0.25
      }
    }

    senao se (tipo == 'I')
    {
      se (consumo <= 100)
      {
        valor = 800
      }
      senao
      {
        valor = 800 + (consumo - 100) * 0.04
      }
    }

    valor = mat.arredondar(valor, 2)

    escreva("CONTA = ", conta, "\n")
    escreva("VALOR DA CONTA = ", valor, "\n")
  }
}