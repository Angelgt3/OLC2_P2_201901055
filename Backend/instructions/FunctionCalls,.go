package instructions

import (
	"Backend/environment"
	"Backend/generator"
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

func (p FunCall) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}
