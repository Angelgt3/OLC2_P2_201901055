package instructions

import (
	"Backend/environment"
)

type FunCall struct {
	Lin        int
	Col        int
	Id         string
	Parametros []interface{}
}

func NewFunCall(lin int, col int, id string, parametros []interface{}) FunCall {
	instr := FunCall{lin, col, id, parametros}
	return instr
}

func (p FunCall) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}
