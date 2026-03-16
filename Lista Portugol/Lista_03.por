programa
{
    funcao inicio()
    {
        inteiro n1, n2, n3
        inteiro numero, quadrado

        escreva("Digite o primeiro número inteiro: ")
        leia(n1)

        escreva("Digite o segundo número inteiro: ")
        leia(n2)

        escreva("Digite o terceiro número inteiro: ")
        leia(n3)

        limpa()

        se (n1 < 0 ou n1 > 9 ou n2 < 0 ou n2 > 9 ou n3 < 0 ou n3 > 9)
        {
            escreva("DÍGITO INVÁLIDO")
        }
        senao
        {
            numero = n1 * 100 + n2 * 10 + n3
            quadrado = numero * numero 

            escreva("O número é: ", numero, "\n")
            escreva("O quadrado do número é: ", quadrado, "\n")
        }
    }
}