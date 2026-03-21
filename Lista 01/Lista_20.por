programa 
{
  funcao inicio() 
  {
inteiro h, m, s
real total_segundos

escreva("Digite a quantidade de horas: ")
leia(h)

escreva("Digite a quantidade de minutos: ")
leia(m)

escreva("Digite a quantidade de segundos: ")
leia(s)

total_segundos = (h * 60 * 60) + (m * 60) + s

escreva("O tempo em segundos é = ", total_segundos, "\n")
  }
}