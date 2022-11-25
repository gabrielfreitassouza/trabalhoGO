package main

import (
	"fmt"
	diretor "trabalho/ByteBank/Funcionarios/Diretor"
	funcionario "trabalho/ByteBank/Funcionarios/Funcionario"
	gerenciador "trabalho/ByteBank/Gereciador"
)

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

	//robertaTeste receber a referncia do tipo diretor
	robertaTeste := roberta

	fmt.Printf("\nBonificacao de uma referencia de Diretor: %.2f", roberta.GetBonificacao())
	fmt.Printf("\nBonificacao de uma referencia de Funcionario: %.2f\n\n", robertaTeste.GetBonificacao())

	gerenciador.Registrar(roberta)

	fmt.Println(carlos.Nome)
	fmt.Printf("%.2f\n\n", carlos.GetBonificacao())

	fmt.Println(roberta.Nome)
	fmt.Printf("%.2f\n\n",roberta.GetBonificacao())

	fmt.Printf("Total de bonificações: %.2f", gerenciador.GetTotalBonificacao())
  
	fmt.Scanf("%d")
}
