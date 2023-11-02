package instructions

import (
	"Backend/environment"
	"Backend/generator"
)

type RemoveLast struct {
	Lin int
	Col int
	Id  string
}

func NewRemoveLast(lin int, col int, id string) RemoveLast {
	instr := RemoveLast{lin, col, id}
	return instr
}

func (p RemoveLast) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INICIO REMOVELAST")
	var result environment.Value

	return result
}
