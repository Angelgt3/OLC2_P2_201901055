package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type If struct {
	Lin         int
	Col         int
	Expresion   interfaces.Expression
	Bloque_if   []interface{}
	Bloque_else []interface{}
	Elif        []interface{}
}

type Elif struct {
	Lin         int
	Col         int
	Expresion   interfaces.Expression
	Bloque_elif []interface{}
}

func NewElif(lin int, col int, condition interfaces.Expression, bloque_elif []interface{}) Elif {
	elifInstr := Elif{lin, col, condition, bloque_elif}
	return elifInstr
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque_if []interface{}, bloque_else []interface{}, ELIF []interface{}) If {
	ifInstr := If{lin, col, condition, bloque_if, bloque_else, ELIF}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}

func (p Elif) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}
