package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type ForRange struct {
	Lin        int
	Col        int
	Id         string
	Expresion  interfaces.Expression
	Expresion2 interfaces.Expression
	Bloque     []interface{}
}

func NewForRange(lin int, col int, id string, condition interfaces.Expression, condition2 interfaces.Expression, bloque []interface{}) ForRange {
	forInstr := ForRange{lin, col, id, condition, condition2, bloque}
	return forInstr
}

func (p ForRange) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}

type For struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewFor(lin int, col int, id string, condition interfaces.Expression, bloque []interface{}) For {
	forInstr := For{lin, col, id, condition, bloque}
	return forInstr
}

func (p For) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}
