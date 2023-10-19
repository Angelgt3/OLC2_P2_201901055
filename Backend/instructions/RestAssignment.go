package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type RestAssigment struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewResAssigment(lin int, col int, id string, val interfaces.Expression) RestAssigment {
	instr := RestAssigment{lin, col, id, val}
	return instr
}

func (p RestAssigment) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}
