package expressions

import (
	"Backend/environment"
	"Backend/generator"
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

func (p CastingInt) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value

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

func (p CastingFloat) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	result := p.Expresion.Ejecutar(ast, env, gen)
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

func (p CastingString) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	result := p.Expresion.Ejecutar(ast, env, gen)
	gen.NativeFloatToString()
	paramTemp := gen.NewTemp()
	gen.AddComment("Empieza el envio de parametros")
	gen.AddExpression(paramTemp, "P", environment.Environment.GetSizeEntorno(env.(environment.Environment)), "+")
	gen.AddSetStack(paramTemp, result.Value)
	gen.AddComment("Termina el envio de parametros")
	gen.AddCall("FloatString")
	temp := gen.NewTemp()
	gen.AddGetStack(temp, "P")

	return result
}
