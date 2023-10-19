package interfaces

import (
	"Backend/environment"
	"Backend/generator"
)

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{}
}
