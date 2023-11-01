package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strings"
)

type While struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, bloque []interface{}) While {
	whileInstr := While{lin, col, condition, bloque}
	return whileInstr
}

func (p While) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("INCIO WHILE")
	//entorno while
	var envwhile environment.Environment
	envwhile = environment.NewEnvironment(env.(environment.Environment), "WHILE", environment.NULL)
	envwhile.Size["size"] = env.(environment.Environment).Size["size"] + 1

	var condicion, result environment.Value
	var OutLvls []interface{}
	inicio := gen.NewLabel()
	gen.AddLabel(inicio)
	condicion = p.Expresion.Ejecutar(ast, envwhile, gen)
	newLabel := gen.NewLabel()
	salida := gen.NewLabel()
	envwhile.ContinueLbl = inicio
	envwhile.BreakLbl = salida

	//add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones
	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			respuesta := s.(interfaces.Instruction).Ejecutar(ast, envwhile, gen)
			if respuesta != nil {
				for _, lvl := range respuesta.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		}
	}
	gen.AddGoto(inicio)

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

	gen.AddLabel(salida)
	gen.AddComment("FIN WHILE")
	return result
}
