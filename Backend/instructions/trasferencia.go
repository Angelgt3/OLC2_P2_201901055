package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type Break struct {
	Lin int
	Col int
}

func NewBreak(lin int, col int) Break {
	Instr := Break{lin, col}
	return Instr
}

func (p Break) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	return "break"
}

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	Instr := Continue{lin, col}
	return Instr
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	return "continue"
}

type Return struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewReturn(lin int, col int, expr interfaces.Expression) Return {
	Instr := Return{lin, col, expr}
	return Instr
}

func (p Return) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	return p.Expresion
}
