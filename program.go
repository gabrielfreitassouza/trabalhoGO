package main

import (
	"fmt"
	"math"
	diretor "trabalho/ByteBank/Funcionarios/Diretor"
	funcionario "trabalho/ByteBank/Funcionarios/Funcionario"
	gerenciador "trabalho/ByteBank/Gereciador"
)

// roundFloat define as casas decimais do valor float
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	//carlos é declarador como tipo Funcionario
	carlos := funcionario.Funcionario{}
	carlos.Nome = "Carlos"
	carlos.CPF = "546.879.157-20"
	carlos.Salario = 2000

	//Executar o função registrar do package gerenciador
	gerenciador.Registrar(carlos)

	//roberta é declador como tipo Diretor
	roberta := diretor.Diretor{}
	roberta.Nome = "Roberta"
	roberta.CPF = "454.658.148-3"
	roberta.Salario = 5000

	//robertaTeste receber  roberta do tipo diretor
	robertaTeste := roberta

	fmt.Println("Bonificacao de uma referencia de Diretor:", roundFloat(roberta.GetBonificacao(), 2))
	fmt.Println("Bonificacao de uma referencia de Funcionario:", roundFloat(robertaTeste.GetBonificacao(), 2))

	gerenciador.Registrar(roberta)

	fmt.Println(carlos.Nome)
	fmt.Println(roundFloat(carlos.GetBonificacao(), 2))

	fmt.Println(roberta.Nome)
	fmt.Println(roundFloat(roberta.GetBonificacao(), 2))

	fmt.Println("Total de bonificações:", roundFloat(gerenciador.GetTotalBonificacao(), 2))
	fmt.Scanf("%d")
}
