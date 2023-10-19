package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type Guard struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewGuard(lin int, col int, condition interfaces.Expression, bloque []interface{}) Guard {
	ifInstr := Guard{lin, col, condition, bloque}
	return ifInstr
}

func (p Guard) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}
