package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
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

func (p CallVar) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("Llamando una variable")
	var result environment.Value
	retSym := env.(environment.Environment).GetVariable(p.Id)
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

func (p FunCallE) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var result environment.Value
	return result
}
