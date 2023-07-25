package main

import (
	"fmt"
)

type Emprestimo struct { //Criando a estrutura do emprestimo
	Score             int
	Renda             float64
	ValorDoEmprestimo float64
	TempoDoEmprestimo int
}

func analise(emprestimo Emprestimo) bool { //funcao responsavel por fazer analise do emprestimo
	Score := emprestimo.Score
	Renda := emprestimo.Renda
	ValorDoEmprestimo := emprestimo.ValorDoEmprestimo
	TempoDoEmprestimo := emprestimo.TempoDoEmprestimo //pegando os valores do emprestimo
	taxa := 0                                         //essa variavel sera atualizada de acordo com o score
	situacao := "APROVADO"                            //situacao do emprestimo e setada como "aprovado", mas pode mudar de acordo com condicionais futuras

	if Score >= 501 {
		taxa = 15
	} else {
		taxa = 20
	} //vai setar o valor da taxa de acordo com o score

	valorParcelaFunc := func(ValorDoEmprestimo float64, TempoDoEmprestimo, taxa int) float64 { //função para pegar o valor das parcelas

		taxaPerc := 0.0 //esta variavel servirá para fazer a multiplicação com valorXtaxa e valorXtempo

		if taxa == 15 { //setando o valor da variavel de acordo com a taxa
			taxaPerc = 0.15
		} else {
			taxaPerc = 0.20
		}

		valorXtaxa := (ValorDoEmprestimo * taxaPerc) / float64(TempoDoEmprestimo)
		valorXtempo := ValorDoEmprestimo / float64(TempoDoEmprestimo) //fazendo os calculos dados pelo enunciado

		valorFinal := float64(valorXtaxa + valorXtempo)

		return valorFinal
	}

	custoEfetivoFunc := func(valorDaParcela, ValorDoEmprestimo float64, TempoDoEmprestimo int) float64 { //função para pegar o custo efetivo
		custoEfetivo := (valorDaParcela * float64(TempoDoEmprestimo)) - ValorDoEmprestimo //fazendo os calculos dados pelo enunciado

		return float64(custoEfetivo)
	}

	valorParcela := valorParcelaFunc(ValorDoEmprestimo, TempoDoEmprestimo, taxa)
	custoEfetivo := custoEfetivoFunc(valorParcela, ValorDoEmprestimo, TempoDoEmprestimo) //setando os valores com as funções

	if Score >= 501 {
		if valorParcela > (Renda * 0.2) { //condicional que vai verificar se o valor da parcela é maior que 20% da renda
			situacao = "RECUSADO"
		}
	} else {
		if valorParcela > (Renda * 0.1) { //entrará neste if caso o socre seja abaixo de 501, ou seja, caso o valor da parcela seja maior que 10%, a situação ficará como "RECUSADO"
			situacao = "RECUSADO"
		}
	}

	if (TempoDoEmprestimo % 12) != 0 { //verifica se o tempo do emprestimo é divisivel por 12, caso não seja erá printar erro no terminal
		fmt.Println("Tempo inválido")
		situacao = "RECUSADO" //e mudar a situação para recusado
	}

	drawTable(Score, TempoDoEmprestimo, taxa, Renda, ValorDoEmprestimo, valorParcela, custoEfetivo, situacao) //chama a função que vai printar a tabela

	return true
}

func drawTable(Score, TempoDoEmprestimo, Taxa int, Renda, ValorDoEmprestimo, ValorParcela, CustoEfetivo float64, Situacao string) string {
	fmt.Println("Análise de crédito para empréstimo")
	fmt.Println("----------------------------------")
	fmt.Printf("Score de crédito:           %v\n", Score)
	fmt.Printf("Renda:                      %v\n", Renda)
	fmt.Printf("Valor do empréstimo:        %v\n", ValorDoEmprestimo)
	fmt.Printf("Tempo do empréstimo:        %v\n", TempoDoEmprestimo)
	fmt.Printf("Valor mensal de parcela:    %2.f\n", ValorParcela)
	fmt.Printf("Taxa de juros:              %v\n", Taxa)
	fmt.Printf("Custo efetivo:              %v\n", CustoEfetivo)
	fmt.Printf("Situação do empréstimo:     %v\n", Situacao)
	fmt.Println("")

	return "" //printando os valores
}

func main() {
	//Instanciando emprestimos para análise e, após, exibição
	var emprestimo1 Emprestimo
	emprestimo1 = Emprestimo{550, 1500.0, 1000.0, 12}

	var emprestimo2 Emprestimo
	emprestimo2 = Emprestimo{501, 1000.0, 1000.0, 24}

	var emprestimo3 Emprestimo
	emprestimo3 = Emprestimo{350, 1000.0, 10000.0, 12}

	analise(emprestimo1)
	analise(emprestimo2)
	analise(emprestimo3)
}
