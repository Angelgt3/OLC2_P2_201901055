package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type StructExp struct {
	Lin     int
	Col     int
	Id      string
	ListExp []interface{}
}

func NewStructExp(lin int, col int, id string, list []interface{}) StructExp {
	exp := StructExp{lin, col, id, list}
	return exp
}

func (p StructExp) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

type StructAccess struct {
	Lin    int
	Col    int
	Struct interfaces.Expression
	Id     string
}

func NewStructAccess(lin int, col int, str interfaces.Expression, id string) StructAccess {
	exp := StructAccess{lin, col, str, id}
	return exp
}

func (p StructAccess) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var result environment.Value
	return result
}
