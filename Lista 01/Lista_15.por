programa 
{
  inclua biblioteca Matematica --> mat 

inteiro n, i  //número digitado
inteiro quadrado // caixa para guardar o resultado da multiplicação

  funcao inicio() 
  {
    escreva("Digite um número inteiro: ")
    leia(n)

    limpa()

    para (i = 1; i <= n; i++) // Isso é um contador que começa a contar de 1 até n
    
    se (i % 2 == 0) // Isso vai verificar vai dividir por 2 e verficar se o resto é zero
    {
      quadrado = i * mat.potencia(i, 2)  // Se for par, isso eleva ele ao quadro.
      escreva(i, " ^ 2 = ", quadrado, "\n") // Isso vai mostrar a saída como se pede.
    }
  }
}
