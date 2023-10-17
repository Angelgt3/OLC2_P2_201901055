package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type Switch struct {
	Lin            int
	Col            int
	Expresion      interfaces.Expression
	Cases          []interface{}
	Bloque_default []interface{}
}

type Cases struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewCase(lin int, col int, condition interfaces.Expression, bloque []interface{}) Cases {
	casesInstr := Cases{lin, col, condition, bloque}
	return casesInstr
}

func NewSwitch(lin int, col int, condition interfaces.Expression, cases []interface{}, bloque_default []interface{}) Switch {
	switchInstr := Switch{lin, col, condition, cases, bloque_default}
	return switchInstr
}

func (p Switch) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	return nil
}

func (p Cases) Ejecutar(ast *environment.AST, env interface{}, condicion environment.Symbol) interface{} {

	return false
}
