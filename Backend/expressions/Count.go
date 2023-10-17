package expressions

import (
	"Backend/environment"
)

type Count struct {
	Lin int
	Col int
	Id  string
}

func NewCount(lin int, col int, id string) Count {
	instr := Count{lin, col, id}
	return instr
}

func (p Count) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: nil}
}
