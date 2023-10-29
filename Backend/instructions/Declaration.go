package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strconv"
	"strings"
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
	var newVar environment.Symbol
	result := p.Expresion.Ejecutar(ast, env, gen)
	// var valor : string ? | variable con tipo sin valor | para eso nomas
	if result.Type == environment.NULL && p.Tipo != environment.NULL {
		result.Type = p.Tipo
	}
	//casteo
	//float | int -> float  | unicamente
	if p.Tipo == environment.FLOAT && result.Type == environment.INTEGER {
		// Convertir el int a float64
		result.Type = environment.FLOAT
	}
	//validar tipos
	if result.Type == environment.ARRAY {
		if p.ArrayValidation(ast, env, result) { //para array

		} else if p.MatrizValidation(ast, env, result) { //para matriz

		} else {
			//ast.SetError("La estructura del array es incorrecta", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
		}
		// variable con el mismo tipo y valor || variable sin tipo asignado pero con valor
	} else if result.Type == p.Tipo || (result.Type != p.Tipo && p.Tipo == environment.NULL) {
		gen.AddComment("Agregando una declaracion variable")
		newVar = env.(environment.Environment).SaveVariable(p.Id, p.changeable, result.Type)
		//ast.SetRs(p.Id, "variable", result.Tipo, env.(environment.Environment).GetEntorno(), p.Lin, p.Col)
	} else {

		//ast.SetError("Los tipos de datos son incorrectos", p.Col, p.Lin, env.(environment.Environment).GetEntorno())
	}

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
	var result environment.Value
	gen.SetMainFlag(false)
	gen.AddComment("FUNCION " + p.Id)
	gen.AddTittle(p.Id)
	//entorno
	var envFunc environment.Environment
	envFunc = environment.NewEnvironment(env.(environment.Environment), p.Id)
	envFunc.Size["size"] = envFunc.Size["size"] + 1
	//variables
	for _, s := range p.Parametros {
		res := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
		envFunc.SaveVariable(res.(environment.Value).Value, true, res.(environment.Value).Type)
	}
	//instrucciones funcion
	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, envFunc, gen)
			if resInst != nil {
				for _, lvl := range resInst.(environment.Value).OutLabel {
					gen.AddLabel(lvl.(string))
				}
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Ejecutar(ast, envFunc, gen)
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}
	gen.AddEnd()
	gen.SetMainFlag(true)
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

func (p Declaration) ArrayValidation(ast *environment.AST, env interface{}, result environment.Value) bool {
	//validaciones de array

	return false
}
func (p Declaration) MatrizValidation(ast *environment.AST, env interface{}, result environment.Value) bool {
	//validaciones de matriz

	return false
}
