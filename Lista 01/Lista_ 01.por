programa 
{ 
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
    real n1, n2, n3, media

    escreva("Digite a primeira nota: ")
    leia(n1)

    escreva("Digite a segunda nota: ")
    leia(n2)

    escreva("Digite a terceira nota: ")
    leia(n3)

    media = (n1 + n2 + n3) / 3.0
    media = mat.arredondar(media, 2)

    se (media >= 6.0)
    {
      escreva("Resultado: Aprovado\n")
    }
    senao
    {
      escreva("Resultado: Reprovado\n")
    }
  }
}
