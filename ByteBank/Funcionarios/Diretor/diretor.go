package diretor

import funcionario "trabalho/ByteBank/Funcionarios/Funcionario"

// Diretor receber o tipo Funcionario
type Diretor funcionario.Funcionario

/*
   Devidor a GO não ter conceito de herança,
   apesar de eu ter definido que Diretor e igual a
   Funcionario, ele  não hedar seus metados. Então
   é necessario recriar ele.
*/

func (d Diretor) GetBonificacao() float64 {
	d.Salario += d.Salario * 0.10
	return d.Salario
}
