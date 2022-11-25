package gerenciador

//totalBonificacao é tipo private
/* Em GO para definir tipo privado basta começar ele
   com letra minuscula. 
*/
var totalBonificacao float64

// Definir o tipo generico da função GetBonificacao() que está em Diretor e Funcionario
type f interface {
	GetBonificacao() float64
}

// Registrar soma o total de bonificações e receber Diretor e Funcionario
func Registrar(fuc f) {
	totalBonificacao += f.GetBonificacao(fuc)
}

// GetTotalBonificacao retorna o valor de totalBonificacao
func GetTotalBonificacao() float64 {
	return totalBonificacao
}
