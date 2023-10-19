package interfaces

import (
	"Backend/environment"
	"Backend/generator"
)

type Expression interface {
	Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value
}
