package expressions

import (
	"Backend/environment"
	"Backend/generator"
)

type IsEmpty struct {
	Lin int
	Col int
	Id  string
}

func NewIsEmpty(lin int, col int, id string) IsEmpty {
	instr := IsEmpty{lin, col, id}
	return instr
}

func (p IsEmpty) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	return environment.Value{}
}
