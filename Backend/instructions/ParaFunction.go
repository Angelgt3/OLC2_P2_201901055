package instructions

import (
	"Backend/environment"
	"Backend/generator"
)

type ParamsDeclaration struct {
	Lin        int
	Col        int
	Id_Externo string
	Id_Interno string
	Inout      string
	Tipo       environment.TipoExpresion
}

func NewParamsDeclaration(lin int, col int, id string, id2 string, inout string, tipo environment.TipoExpresion) ParamsDeclaration {
	instr := ParamsDeclaration{lin, col, id, id2, inout, tipo}
	return instr
}

func (p ParamsDeclaration) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

	return nil
}
