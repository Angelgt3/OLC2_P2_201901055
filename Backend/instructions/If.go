package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strings"
)

type If struct {
	Lin         int
	Col         int
	Expresion   interfaces.Expression
	Bloque_if   []interface{}
	Bloque_else []interface{}
	Elif        []interface{}
}

type Elif struct {
	Lin         int
	Col         int
	Expresion   interfaces.Expression
	Bloque_elif []interface{}
}

func NewElif(lin int, col int, condition interfaces.Expression, bloque_elif []interface{}) Elif {
	elifInstr := Elif{lin, col, condition, bloque_elif}
	return elifInstr
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque_if []interface{}, bloque_else []interface{}, ELIF []interface{}) If {
	ifInstr := If{lin, col, condition, bloque_if, bloque_else, ELIF}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INICIO IF")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Expresion.Ejecutar(ast, env, gen)
	newLabel := gen.NewLabel()
	salida := gen.NewLabel()
	//add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones if
	for _, s := range p.Bloque_if {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			//ejecutar instrucciones
			respuesta := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if respuesta != nil {
				//agregando etiquetas de salida
				for _, lvl := range respuesta.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		}
	}
	//add out labels
	gen.AddGoto(salida)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones elif
	if p.Elif != nil {
		for _, ins := range p.Elif {
			if elifInstance, ok := ins.(Elif); ok {
				//ejecutar instrucciones
				respuesta := elifInstance.Ejecutar(ast, env, gen, salida)
				if respuesta != nil {
					//agregando etiquetas de salida
					for _, lvl := range respuesta.(environment.Value).OutLabel {
						OutLvls = append(OutLvls, lvl)
					}
				}
			}
		}
	}

	//instrucciones else
	for _, s := range p.Bloque_else {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			//ejecutar instrucciones
			respuesta := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if respuesta != nil {
				//agregando etiquetas de salida
				for _, lvl := range respuesta.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		}
	}

	OutLvls = append(OutLvls, newLabel)
	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}
	result.OutLabel = copiedSlice

	gen.AddLabel(salida)
	gen.AddComment("FIN If")
	return result
}

func (p Elif) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator, salida string) interface{} {
	gen.AddComment("INICIO ELIf")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Expresion.Ejecutar(ast, env, gen)
	newLabel := gen.NewLabel()
	//add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones elif
	for _, s := range p.Bloque_elif {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			respuesta := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if respuesta != nil {
				//agregando etiquetas de salida
				for _, lvl := range respuesta.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		}
	}
	//add out
	gen.AddGoto(salida)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	OutLvls = append(OutLvls, newLabel)
	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	gen.AddComment("FIN ELIF")
	return result
}
