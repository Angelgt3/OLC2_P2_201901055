package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

// variables
type Declaration struct {
	Lin        int
	Col        int
	Id         string
	Tipo       environment.TipoExpresion
	Expresion  interfaces.Expression
	changeable bool
}

func NewDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, val interfaces.Expression, mut bool) Declaration {
	instr := Declaration{lin, col, id, tipo, val, mut}
	return instr
}

func (p Declaration) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}

// funciones
type DeclarationFunc struct {
	Lin        int
	Col        int
	Id         string
	Tipo       environment.TipoExpresion
	Parametros []interface{}
	Bloque     []interface{}
}

func NewDeclarationFunc(lin int, col int, id string, tipo environment.TipoExpresion, parametros []interface{}, bloque []interface{}) DeclarationFunc {
	instr := DeclarationFunc{lin, col, id, tipo, parametros, bloque}
	return instr
}

func (p DeclarationFunc) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}

// Declaracion de struct
type Struct struct {
	Lin     int
	Col     int
	Id      string
	ListAtr []interface{}
}

func NewStruct(lin int, col int, id string, list []interface{}) Struct {
	instr := Struct{lin, col, id, list}
	return instr
}

func (p Struct) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	env.(environment.Environment).SaveStruct(p.Id, p.ListAtr)
	return nil
}
