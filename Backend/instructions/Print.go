package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
)

type Print struct {
	Lin   int
	Col   int
	Value interface{}
}

func NewPrint(lin int, col int, val interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result = p.Value.(interfaces.Expression).Ejecutar(ast, env, gen)
	if result.Type == environment.INTEGER || result.Type == environment.FLOAT || result.Type == environment.ARRAY {
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
		gen.AddPrintf("c", "10")
		gen.AddBr()
	}
	return nil
}
