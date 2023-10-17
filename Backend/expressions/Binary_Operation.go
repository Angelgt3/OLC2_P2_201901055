package expressions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type BinaryOperation struct {
	Lin      int
	Col      int
	Op_izq   interfaces.Expression
	Operador string
	Op_der   interfaces.Expression
}

func NewOperationBinary(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) BinaryOperation {
	exp := BinaryOperation{Lin: lin, Col: col, Op_izq: Op1, Operador: Operador, Op_der: Op2}
	return exp
}

func (operacion BinaryOperation) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	return environment.Symbol{Lin: operacion.Lin, Col: operacion.Col, Tipo: environment.NULL, Valor: nil}
}
