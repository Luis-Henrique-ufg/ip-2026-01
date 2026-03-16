programa 
{
  inclua biblioteca Matematica --> mat 

  funcao inicio() {
    real salario_minimo, qtd_kw, valor_100kw, valor_kw, custo_consumo, custo_desconto

    escreva("Qual o salário mínimo? ")
    leia(salario_minimo)

    escreva("Qual a quantidade de kW consumida pela residência? ")
    leia(qtd_kw)

    limpa()

    valor_100kw = salario_minimo * 0.70
    valor_kw = valor_100kw / 100
    valor_kw = mat.arredondar(valor_kw, 2)

    custo_consumo = qtd_kw * valor_kw
    custo_consumo = mat.arredondar(custo_consumo, 2)
    custo_desconto = custo_consumo * 0.90
    custo_desconto = mat.arredondar(custo_desconto, 2)

    escreva("O valor em reais de cada kW: R$", valor_kw, "\n")
    escreva("O valor em reais a ser pago pelo consumo da residência: R$ ", custo_consumo, "\n")
    escreva("O novo valor a ser pago pela residência com um desconto de 10%: R$ ", custo_desconto, "\n")
  }
}
