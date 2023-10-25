package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
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

func (p Declaration) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	var newVar environment.Symbol
	result = p.Expresion.Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	newVar = env.(environment.Environment).SaveVariable(p.Id, p.changeable, p.Tipo)

	if result.Type == environment.BOOLEAN {

		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), "1")
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), "0")
		gen.AddGoto(newLabel)
		gen.AddLabel(newLabel)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
		gen.AddBr()
	}

	return result
}

// validacion matriz
func (p Declaration) MatrizValidation(ast *environment.AST, env interface{}, result environment.Symbol) bool {

	return true
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

func (p DeclarationFunc) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {

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

func (p Struct) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	env.(environment.Environment).SaveStruct(p.Id, p.ListAtr)
	return nil
}
