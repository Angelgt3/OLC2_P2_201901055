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
	var result, op1 environment.Value
	if operacion.Operador == "!" {
		op1 = operacion.Op_der.Ejecutar(ast, env, gen)
		if op1.Type == environment.BOOLEAN {
			result = environment.NewValue("", false, environment.BOOLEAN)
			result.TrueLabel = append(op1.FalseLabel, result.TrueLabel...)
			result.FalseLabel = append(op1.TrueLabel, result.FalseLabel...)
			return result
		} else {
			ast.SetError("El tipo no compatible para realizar el NOT", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
	} else if operacion.Operador == "-" {
		newTemp := gen.NewTemp()
		op1 = operacion.Op_der.Ejecutar(ast, env, gen)
		if op1.Type == environment.INTEGER || op1.Type == environment.FLOAT {
			gen.AddExpression(newTemp, op1.Value, "-1", "*")
			result = environment.NewValue(newTemp, true, environment.FLOAT)
			result.IntValue = op1.IntValue + -1
			return result
		} else {
			ast.SetError("El tipo no compatible para realizar la negacion", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
	}

	gen.AddBr()
	return environment.Value{}
}
