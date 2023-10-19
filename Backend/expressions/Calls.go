package expressions

import (
	"Backend/environment"
	"Backend/generator"
)

type CallVar struct {
	Lin int
	Col int
	Id  string
}

func NewCallVariable(lin int, col int, id string) CallVar {
	exp := CallVar{lin, col, id}
	return exp
}

func (p CallVar) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

type FunCallE struct {
	Lin        int
	Col        int
	Id         string
	Parametros []interface{}
}

func NewFunCallE(lin int, col int, id string, parametros []interface{}) FunCallE {
	instr := FunCallE{lin, col, id, parametros}
	return instr
}

func (p FunCallE) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var result environment.Value
	return result
}
