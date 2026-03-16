programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() 
  {
  real h, a
  real area_base, volume

  escreva("Digite a altura da pirâmide: ")  
  leia(h)

  escreva("Digite o tamanho da aresta da base: ")
  leia(a)

  limpa()

  area_base = (3 * mat.potencia(a, 2.0) * mat.raiz(3.0, 2.0)) / 2.0
  volume = (1.0 / 3.0) * area_base * h 
  volume = mat.arredondar(volume, 2)

  escreva ("O voume da pirâmide é: ", volume, " metros cúbicos\n")
  }
}
