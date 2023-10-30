package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type Assigment struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewAssigment(lin int, col int, id string, val interfaces.Expression) Assigment {
	instr := Assigment{lin, col, id, val}
	return instr
}

func (p Assigment) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	gen.AddComment("Generando asignacion")
	variable := env.(environment.Environment).GetVariable(p.Id)
	if !variable.Mutable {
		ast.SetError("La variable es inmutable", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
		return result
	}
	result = p.Expresion.Ejecutar(ast, env, gen)
	gen.AddSetStack(strconv.Itoa(variable.Posicion), result.Value)
	gen.AddBr()
	return result
}

type AssigmentVec struct {
	Lin        int
	Col        int
	Id         string
	Expresion  interfaces.Expression
	Expresion2 interfaces.Expression
}

func NewAssigmentVec(lin int, col int, id string, val interfaces.Expression, val2 interfaces.Expression) AssigmentVec {
	instr := AssigmentVec{lin, col, id, val, val2}
	return instr
}

func (p AssigmentVec) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}

type AssigmentMatriz struct {
	Lin        int
	Col        int
	Id         string
	Expresion  interfaces.Expression
	Expresion2 interfaces.Expression
}

func NewAssigmentMa(lin int, col int, id string, val interfaces.Expression, val2 interfaces.Expression) AssigmentMatriz {
	instr := AssigmentMatriz{lin, col, id, val, val2}
	return instr
}

func (p AssigmentMatriz) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value

	return result
}
