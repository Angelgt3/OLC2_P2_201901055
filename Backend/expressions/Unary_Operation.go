package expressions

import (
	"Backend/environment"
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

func (operacion UnaryOperation) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result interface{}
	return environment.Symbol{Lin: operacion.Lin, Col: operacion.Col, Tipo: environment.NULL, Valor: result}
}
