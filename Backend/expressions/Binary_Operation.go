package expressions

import (
	"Backend/environment"
	"Backend/generator"
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

func (operacion BinaryOperation) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var dominante environment.TipoExpresion

	tabla_dominante := [3][3]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING
		{environment.INTEGER, environment.FLOAT, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.NULL},
		//STRING
		{environment.NULL, environment.NULL, environment.STRING},
	}

	var op1, op2, result environment.Value
	newTemp := gen.NewTemp()

	//SUMA,RESTA,MULTIPLICACION
	if operacion.Operador == "+" || operacion.Operador == "-" || operacion.Operador == "*" {
		//realiza la operaciones
		op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
		op2 = operacion.Op_der.Ejecutar(ast, env, gen)
		dominante = tabla_dominante[op1.Type][op2.Type]
		if dominante == environment.INTEGER || dominante == environment.FLOAT { //valida el tipo
			gen.AddExpression(newTemp, op1.Value, op2.Value, operacion.Operador)
			result = environment.NewValue(newTemp, true, dominante)
			result.IntValue = op1.IntValue + op2.IntValue
		} else {
			ast.SetError(" No es posible realizar la operaciones '"+operacion.Operador+"'", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
	} else if operacion.Operador == "/" { //DIVISION
		/*
			op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
			op2 = operacion.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante)
			}
		*/
	}

	switch operacion.Operador {

	case "%":
		{

		}
	case "<":
		{

		}
	case ">":
		{

		}
	case "<=":
		{

		}
	case ">=":
		{

		}
	case "==":
		{

		}
	case "!=":
		{

		}
	case "&&":
		{

		}
	case "||":

	}
	return result
}
