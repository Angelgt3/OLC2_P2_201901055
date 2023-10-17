package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type Append struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewAppend(lin int, col int, id string, val interfaces.Expression) Append {
	instr := Append{lin, col, id, val}
	return instr
}

func (p Append) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	//Traer simbolo
	return nil
}
