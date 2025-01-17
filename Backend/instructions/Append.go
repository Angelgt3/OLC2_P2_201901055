package instructions

import (
	"Backend/environment"
	"Backend/generator"
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

func (a Append) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("CREANDO APPEND")

	var result environment.Value

	return result
}
