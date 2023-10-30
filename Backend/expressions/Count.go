package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type Count struct {
	Lin int
	Col int
	Id  string
}

func NewCount(lin int, col int, id string) Count {
	instr := Count{lin, col, id}
	return instr
}

func (p Count) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var result environment.Value
	variable := env.(environment.Environment).GetVariable(p.Id)

	tem := gen.NewTemp()
	gen.AddGetStack(tem, strconv.Itoa(variable.Posicion))

	gen.AddBr()
	return result
}
