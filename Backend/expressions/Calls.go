package expressions

import (
	"Backend/environment"
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

func (p CallVar) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Environment).GetVariable(p.Id)
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

func (p FunCallE) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result environment.Symbol
	return result
}
