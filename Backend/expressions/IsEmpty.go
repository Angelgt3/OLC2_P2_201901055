package expressions

import (
	"Backend/environment"
)

type IsEmpty struct {
	Lin int
	Col int
	Id  string
}

func NewIsEmpty(lin int, col int, id string) IsEmpty {
	instr := IsEmpty{lin, col, id}
	return instr
}

func (p IsEmpty) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: nil}
}
