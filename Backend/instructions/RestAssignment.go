package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type RestAssigment struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewResAssigment(lin int, col int, id string, val interfaces.Expression) RestAssigment {
	instr := RestAssigment{lin, col, id, val}
	return instr
}

func (p RestAssigment) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	//Traer simbolo
	result := p.Expresion.Ejecutar(ast, env, gen)

	//Obtengo la variable
	var variable = environment.Symbol{}
	variable = env.(environment.Environment).GetVariable(p.Id)

	//hago la operacion
	if variable.Tipo == environment.FLOAT && result.Type == environment.FLOAT {
		result.Type = environment.FLOAT
	} else if variable.Tipo == environment.FLOAT && result.Type == environment.INTEGER {
		result.Type = environment.FLOAT
	} else if variable.Tipo == environment.INTEGER && result.Type == environment.INTEGER {
		result.Type = environment.INTEGER
	} else {
		ast.SetError(" No es posible realizar la restar", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
		return result
	}
	gen.AddComment("DECREMENTO")
	//modificar la variable si se puede
	newTemp := gen.NewTemp()
	gen.AddGetStack(newTemp, strconv.Itoa(variable.Posicion))
	gen.AddExpression(newTemp, newTemp, result.Value, "-") //suma
	//result.IntValue =  + result.IntValue
	gen.AddSetStack(strconv.Itoa(variable.Posicion), newTemp)
	gen.AddBr()

	return result
}
