package expressions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type CastingInt struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewCastingInt(lin int, col int, val interfaces.Expression) CastingInt {
	instr := CastingInt{lin, col, val}
	return instr
}

func (p CastingInt) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result = p.Expresion.Ejecutar(ast, env)
	return result
}

type CastingFloat struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewCastingFloat(lin int, col int, val interfaces.Expression) CastingFloat {
	instr := CastingFloat{lin, col, val}
	return instr
}

func (p CastingFloat) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result = p.Expresion.Ejecutar(ast, env)
	return result
}

type CastingString struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewCastingString(lin int, col int, val interfaces.Expression) CastingString {
	instr := CastingString{lin, col, val}
	return instr
}

func (p CastingString) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result = p.Expresion.Ejecutar(ast, env)
	return result
}
