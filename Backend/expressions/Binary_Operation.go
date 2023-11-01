package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
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

	//SUMA,RESTA,MULTIPLICACION,MODULO -- ARITMETICA
	if operacion.Operador == "+" || operacion.Operador == "-" || operacion.Operador == "*" || operacion.Operador == "%" {
		//realiza la operaciones
		op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
		op2 = operacion.Op_der.Ejecutar(ast, env, gen)
		dominante = tabla_dominante[op1.Type][op2.Type]
		if dominante == environment.INTEGER || dominante == environment.FLOAT { //valida el tipo
			if operacion.Operador == "%" {
				gen.AddExpression(newTemp, "(int)"+op1.Value, op2.Value, operacion.Operador)
			} else {
				gen.AddExpression(newTemp, op1.Value, op2.Value, operacion.Operador)
			}
			result = environment.NewValue(newTemp, true, dominante)
			result.IntValue = op1.IntValue + op2.IntValue
			return result
		} else if operacion.Operador == "+" && dominante == environment.STRING {
			gen.NativeConcatString()
			//concat
			envSize := strconv.Itoa(env.(environment.Environment).Size["size"])
			tmp1 := gen.NewTemp()
			tmp2 := gen.NewTemp()
			gen.AddExpression(tmp1, "P", envSize, "+")
			gen.AddExpression(tmp1, tmp1, "1", "+")
			gen.AddSetStack("(int)"+tmp1, op1.Value)
			gen.AddExpression(tmp1, tmp1, "1", "+")
			gen.AddSetStack("(int)"+tmp1, op2.Value)
			gen.AddExpression("P", "P", envSize, "+")
			gen.AddCall("dbrust_concatString")
			gen.AddGetStack(tmp2, "(int)P")
			gen.AddExpression("P", "P", envSize, "-")
			gen.AddBr()
			result = environment.NewValue(tmp2, true, dominante)
			return result
		} else {
			ast.SetError(" No es posible realizar la operacion '"+operacion.Operador+"'", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
	} else if operacion.Operador == "/" { //DIVISION
		op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
		op2 = operacion.Op_der.Ejecutar(ast, env, gen)
		dominante = tabla_dominante[op1.Type][op2.Type]
		if dominante == environment.INTEGER || dominante == environment.FLOAT {
			MathError := gen.NewLabel()
			Terminar := gen.NewLabel()

			gen.AddIf(op2.Value, "0", "!=", MathError)
			gen.PrintMathError()
			gen.AddExpression(newTemp, "0", "", "")
			gen.AddGoto(Terminar)
			gen.AddLabel(MathError)
			gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
			gen.AddLabel(Terminar)
			result = environment.NewValue(newTemp, true, dominante)
			return result
		} else {
			ast.SetError(" No es posible realizar la operacion '"+operacion.Operador+"'", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
		// < , > , <= , >= , == , != -- RELACIONAL
	} else if operacion.Operador == "<" || operacion.Operador == ">" || operacion.Operador == "<=" || operacion.Operador == ">=" || operacion.Operador == "==" || operacion.Operador == "!=" {
		op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
		op2 = operacion.Op_der.Ejecutar(ast, env, gen)
		dominante = tabla_dominante[op1.Type][op2.Type]
		if dominante == environment.INTEGER || dominante == environment.FLOAT {
			trueLabel := gen.NewLabel()
			falseLabel := gen.NewLabel()
			gen.AddIf(op1.Value, op2.Value, operacion.Operador, trueLabel)
			gen.AddGoto(falseLabel)
			result = environment.NewValue("", false, environment.BOOLEAN)
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else {
			ast.SetError(" No es posible realizar la operacion '"+operacion.Operador+"'", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
		// -- LOGICA
	} else if operacion.Operador == "&&" { // &&
		op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
		for _, lvl := range op1.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		op2 = operacion.Op_der.Ejecutar(ast, env, gen)
		if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
			result = environment.NewValue("", false, environment.BOOLEAN)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op1.FalseLabel, result.FalseLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		} else {
			ast.SetError(" No es posible realizar la operacion '"+operacion.Operador+"'", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
	} else if operacion.Operador == "||" { // ||
		op1 = operacion.Op_izq.Ejecutar(ast, env, gen)
		for _, lvl := range op1.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		op2 = operacion.Op_der.Ejecutar(ast, env, gen)
		if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
			result = environment.NewValue("", false, environment.BOOLEAN)
			result.TrueLabel = append(op1.TrueLabel, result.TrueLabel...)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		} else {
			ast.SetError(" No es posible realizar la operacion '"+operacion.Operador+"'", operacion.Col, operacion.Lin, env.(environment.Environment).GetEntorno())
		}
	}
	gen.AddBr()
	return environment.Value{}
}
