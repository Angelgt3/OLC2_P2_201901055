package instructions

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strings"
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

func (p Switch) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generando SWITCH")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Expresion.Ejecutar(ast, env, gen)
	newLabel := gen.NewLabel()
	truelabel := gen.NewLabel()
	falselabel := gen.NewLabel()
	//add true labels
	gen.AddLabel(truelabel)
	//instrucciones cases
	if p.Cases != nil {
		for _, ins := range p.Cases {
			if elifInstance, ok := ins.(Cases); ok {
				resInst := elifInstance.Ejecutar(ast, env, gen, newLabel, condicion)
				if resInst != nil {
					//agregando etiquetas de salida
					for _, lvl := range resInst.(environment.Value).OutLabel {
						OutLvls = append(OutLvls, lvl)
					}
				}
			}
		}
	}
	//add out labels
	gen.AddGoto(falselabel)

	//add false labels
	gen.AddLabel(falselabel)
	//instrucciones DEFAULT
	gen.AddComment("Generando DEFAULT")
	for _, s := range p.Bloque_default {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				//agregando etiquetas de salida
				for _, lvl := range resInst.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}

	OutLvls = append(OutLvls, newLabel)
	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	return result
}

func (p Cases) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator, salida string, condicion1 environment.Value) interface{} {
	gen.AddComment("Generando CASES")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Expresion.Ejecutar(ast, env, gen)
	newLabel := gen.NewLabel()
	//add true labels
	trueLabel := gen.NewLabel()
	falseLabel := gen.NewLabel()
	gen.AddIf(condicion.Value, condicion1.Value, "==", trueLabel)
	gen.AddGoto(falseLabel)
	gen.AddLabel(trueLabel)

	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				//agregando etiquetas de salida
				for _, lvl := range resInst.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}

	//add out labels
	gen.AddGoto(salida)
	gen.AddLabel(falseLabel)
	OutLvls = append(OutLvls, newLabel)
	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	return result
}
