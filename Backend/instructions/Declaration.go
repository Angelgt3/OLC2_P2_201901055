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

func (d Declaration) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var newVar environment.Symbol
	result := d.Expresion.Ejecutar(ast, env, gen)

	// var valor : string ? | variable con tipo sin valor | para eso nomas
	if result.Type == environment.NULL && d.Tipo != environment.NULL {
		result.Type = d.Tipo
	}
	//casteo
	//float | int -> float  | unicamente
	if d.Tipo == environment.FLOAT && result.Type == environment.INTEGER {
		// Convertir el int a float64
		result.Type = environment.FLOAT
	}
	//validar tipos
	if result.Type == environment.ARRAY {
		newVar = env.(environment.Environment).SaveArray(d.Id, d.Tipo, len(result.ArrValue))
		gen.AddComment("Iniciando la declaración de un ARRAY")
		d.ArrayValidation(ast, env, gen, result.ArrValue)
		gen.AddComment("Se finalizó la declaración de un ARRAY")
		return result
		// variable con el mismo tipo y valor || variable sin tipo asignado pero con valor
	} else if result.Type == d.Tipo || (result.Type != d.Tipo && d.Tipo == environment.NULL) {
		gen.AddComment("Agregando una declaracion variable")
		newVar = env.(environment.Environment).SaveVariable(d.Id, d.changeable, result.Type)
		ast.SetRs(d.Id, "variable", result.Type, env.(environment.Environment).GetEntorno(), d.Lin, d.Col)
	} else {
		ast.SetError("Los tipos de datos son incorrectos para realizar la declaracion", d.Col, d.Lin, env.(environment.Environment).GetEntorno())
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

func (d DeclarationFunc) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	gen.SetMainFlag(false)
	gen.AddComment("FUNCION " + d.Id)
	gen.AddTittle(d.Id)
	//guardo la funcion
	env.(environment.Environment).SaveFunc(d.Id, environment.InstF{d.Lin, d.Col, d.Id, d.Tipo})
	//entorno
	envFunc := environment.NewEnvironment(env.(environment.Environment), d.Id, d.Tipo)
	envFunc.Size["size"] = env.(environment.Environment).Size["size"] + 1
	returnlb := gen.NewLabel()
	envFunc.ReturnLbl = returnlb

	//variables
	for _, s := range d.Parametros {
		res := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
		if result, ok := res.(environment.Value); ok {
			envFunc.SaveVariable(result.Value, true, result.Type)
			ast.SetRs(d.Id, "variable", result.Type, env.(environment.Environment).GetEntorno(), d.Lin, d.Col)
		}

	}
	//instrucciones funcion
	for _, s := range d.Bloque {
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
		}
	}
	gen.AddLabel(returnlb)
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

func (d Struct) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	env.(environment.Environment).SaveStruct(d.Id, d.ListAtr)
	return nil
}

func (d Declaration) ArrayValidation(ast *environment.AST, env interface{}, gen *generator.Generator, arreglo []interface{}) {
	//validaciones de vector
	for _, exp := range arreglo {
		if exp.(environment.Value).Type == environment.ARRAY {
			d.ArrayValidation(ast, env, gen, exp.(environment.Value).ArrValue)
		} else {
			envSize := env.(environment.Environment).NewVariable()
			gen.AddSetStack(strconv.Itoa(envSize.Posicion), exp.(environment.Value).Value)
			gen.AddBr()
		}
	}
}
func (d Declaration) MatrizValidation(ast *environment.AST, env interface{}, result environment.Value) bool {
	//validaciones de matriz

	return true
}
