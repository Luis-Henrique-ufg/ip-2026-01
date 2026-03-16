programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
   
   real f, c, p, mm 
   inteiro i 

   para (i = 1; i <= 3; i++)
   {
   escreva("Digite a temperatura em graus Fahrenheit: ")
   leia(f)

   c = 5 * (f - 32) / 9
   c = mat.arredondar(c, 2)

   escreva("Digite a quantidade de chuva em polegadas: ")
   leia(p)
   
   mm = p * 25.4
   mm = mat.arredondar(mm, 2)

   limpa()

  escreva("\nO valor em graus Celsisus é = ", c, " °C\n") 
  escreva("\nA quantidade de chuva é = ", mm, " mm\n\n")
   }

  }
}
