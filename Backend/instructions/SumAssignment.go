package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type SumAssigment struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewSumAssigment(lin int, col int, id string, val interfaces.Expression) SumAssigment {
	instr := SumAssigment{lin, col, id, val}
	return instr
}

func (p SumAssigment) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}
