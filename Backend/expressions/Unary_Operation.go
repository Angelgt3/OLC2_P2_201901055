package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type UnaryOperation struct {
	Lin      int
	Col      int
	Op_der   interfaces.Expression
	Operador string
}

func NewOperationUnary(lin int, col int, Op1 interfaces.Expression, Operador string) UnaryOperation {
	exp := UnaryOperation{Lin: lin, Col: col, Op_der: Op1, Operador: Operador}
	return exp
}

func (operacion UnaryOperation) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	//var op1 = operacion.Op_der.Ejecutar(ast, env, gen)

	switch operacion.Operador {
	case "!":
		{

		}
	case "-":
		{

		}
	}

	return result
}
