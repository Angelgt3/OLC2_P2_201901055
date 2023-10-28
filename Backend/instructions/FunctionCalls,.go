package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type FunCall struct {
	Lin        int
	Col        int
	Id         string
	Parametros []interface{}
}

func NewFunCall(lin int, col int, id string, parametros []interface{}) FunCall {
	instr := FunCall{lin, col, id, parametros}
	return instr
}

func (fc FunCall) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	size := env.(environment.Environment).Size["size"]
	gen.AddComment("Inicio Llamando Funcion (expresion)")
	if len(fc.Parametros) > 0 {
		tmp1 := gen.NewTemp()
		gen.AddExpression(tmp1, "P", strconv.Itoa(size+1), "+")
		for i := 0; i < len(fc.Parametros); i++ { //recorrer la lista de parametros
			if strings.Contains(fmt.Sprintf("%T", fc.Parametros[i]), "expressions") { //ejecuta la expresion
				result = fc.Parametros[i].(interfaces.Expression).Ejecutar(ast, env, gen)
			}
			//guardando
			gen.AddSetStack("(int)"+tmp1, result.Value)
			if (len(fc.Parametros) - 1) != i {
				gen.AddExpression(tmp1, tmp1, "1", "+")
			}
		}
		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(fc.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")
	} else {
		tmp1 := gen.NewTemp()
		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(fc.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")
	}
	gen.AddComment("Final de llamada funcion (expreision)")
	return result
}
