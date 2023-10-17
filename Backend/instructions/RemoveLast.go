package instructions

import (
	"Backend/environment"
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

func (p RemoveLast) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}
