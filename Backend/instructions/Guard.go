package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strings"
)

type Guard struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewGuard(lin int, col int, condition interfaces.Expression, bloque []interface{}) Guard {
	ifInstr := Guard{lin, col, condition, bloque}
	return ifInstr
}

func (p Guard) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INICIO GUARD")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Expresion.Ejecutar(ast, env, gen)
	newLabel := gen.NewLabel()
	salida := gen.NewLabel()
	//add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//add out labels
	gen.AddGoto(salida)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	//instrucciones guard
	for _, s := range p.Bloque {
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
	OutLvls = append(OutLvls, newLabel)
	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}
	result.OutLabel = copiedSlice
	gen.AddLabel(salida)
	gen.AddComment("FIN GUARD")
	return result
}
