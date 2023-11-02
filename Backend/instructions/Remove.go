package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type Remove struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewRemove(lin int, col int, id string, expr interfaces.Expression) Remove {
	instr := Remove{lin, col, id, expr}
	return instr
}

func (p Remove) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INICIO REMOVE")
	var result environment.Value

	return result
}
