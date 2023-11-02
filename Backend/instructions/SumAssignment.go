package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type SumAssigment struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewSumAssigment(lin int, col int, id string, val interfaces.Expression) SumAssigment {
	instr := SumAssigment{lin, col, id, val}
	return instr
}

func (p SumAssigment) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INICIO INCREMENTO")
	//Traer simbolo
	var result environment.Value
	result = p.Expresion.Ejecutar(ast, env, gen)

	//Obtengo la variable
	variable := env.(environment.Environment).GetVariable(p.Id)

	//hago la operacion
	if variable.Tipo == environment.FLOAT && result.Type == environment.FLOAT {
		variable.Tipo = environment.FLOAT
	} else if variable.Tipo == environment.FLOAT && result.Type == environment.INTEGER {
		variable.Tipo = environment.FLOAT
	} else if variable.Tipo == environment.INTEGER && result.Type == environment.INTEGER {
		variable.Tipo = environment.INTEGER
		/*} else if variable.Tipo == environment.STRING && result.Tipo == environment.STRING {
		result.Valor = strconv.Itoa(variable.Valor.(int)) + strconv.Itoa(result.Valor.(int))
		result.Tipo = environment.STRING
		*/
	} else {
		ast.SetError("No se puedo realizar la suma", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
		return result
	}
	//modificar la variable si se puede
	newTemp := gen.NewTemp()
	gen.AddGetStack(newTemp, strconv.Itoa(variable.Posicion))
	gen.AddExpression(newTemp, newTemp, result.Value, "+") //suma
	//result.IntValue =  + result.IntValue
	gen.AddSetStack(strconv.Itoa(variable.Posicion), newTemp)
	gen.AddComment("FIN INCREMENTO")
	gen.AddBr()

	return result
}
