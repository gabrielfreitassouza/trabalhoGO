package funcionario

/*
		Em GO não liguagem orientada a objetos classica,
	    então não tem o conceito de classe e nem de herança.
	    Porem tem algo similar que são os struct que podem
	    ter metados e interfaces.
*/
type Funcionario struct {
	Nome    string
	CPF     string
	Salario float64
}

// GetBonificacao é metado que retorna a bonificacão do Funcionario
func (f Funcionario) GetBonificacao() float64 {
	return f.Salario * 0.10
}

// f é interface do metado GetBonificacao
type f interface {
	GetBonificacao() float64
}
