package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type CallVar struct {
	Lin int
	Col int
	Id  string
}

func NewCallVariable(lin int, col int, id string) CallVar {
	exp := CallVar{lin, col, id}
	return exp
}

func (c CallVar) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("Llamando una variable")
	var result environment.Value
	retSym := env.(environment.Environment).GetVariable(c.Id)
	newTemp := gen.NewTemp()
	newTemp2 := gen.NewTemp()
	if gen.MainCode {
		gen.AddGetStack(newTemp2, strconv.Itoa(retSym.Posicion))
	} else {
		gen.AddExpression(newTemp, "P", strconv.Itoa(retSym.Posicion), "+")
		gen.AddGetStack(newTemp2, "(int)"+newTemp)
	}

	if retSym.Tipo == environment.BOOLEAN {
		trueLabel := gen.NewLabel()
		falseLabel := gen.NewLabel()
		gen.AddIf(newTemp2, "1", "==", trueLabel)
		gen.AddGoto(falseLabel)
		result = environment.NewValue("", false, environment.BOOLEAN)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	} else {
		result = environment.NewValue(newTemp2, true, retSym.Tipo)
	}
	return result
}

type FunCallE struct {
	Lin        int
	Col        int
	Id         string
	Parametros []interface{}
}

func NewFunCallE(lin int, col int, id string, parametros []interface{}) FunCallE {
	instr := FunCallE{lin, col, id, parametros}
	return instr
}

func (c FunCallE) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	size := env.(environment.Environment).Size["size"]
	gen.AddComment("Llamando Funcion")
	if len(c.Parametros) > 0 {
		tmp1 := gen.NewTemp()
		gen.AddExpression(tmp1, "P", strconv.Itoa(size+1), "+")
		for i := 0; i < len(c.Parametros); i++ { //recorrer la lista de parametros
			if strings.Contains(fmt.Sprintf("%T", c.Parametros[i]), "expressions") { //ejecuta la expresion
				result = c.Parametros[i].(interfaces.Expression).Ejecutar(ast, env, gen)
			}
			//guardando
			gen.AddSetStack("(int)"+tmp1, result.Value)
			if (len(c.Parametros) - 1) != i {
				gen.AddExpression(tmp1, tmp1, "1", "+")
			}
		}
		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(c.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")
	} else {
		tmp1 := gen.NewTemp()
		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(c.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")
	}
	gen.AddComment("Final de llamada funcion")
	return result
}
