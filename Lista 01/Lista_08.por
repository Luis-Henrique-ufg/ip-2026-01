programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
    real r, h
    real a_t, a_c, a_l

    escreva("Qual o raio da lata? ")
    leia(r)

    escreva("Qual a altura da lata? ")
    leia(h)

    a_c = 3.14159 * mat.potencia(r, 2)
    a_l = 2 * 3.14159 * r * h
    a_t = 2 * a_c + a_l

    a_t = a_t * 100
    a_t = mat.arredondar(a_t, 2)

    escreva("\nO valor do custo da lata é: ", a_t, " reais\n")
  }
}
